package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Package struct {
	Id         int       `orm:"column(id);auto"`
	OwnerId    *User     `orm:"column(owner_id);rel(fk)"`
	ReceiverId *User     `orm:"column(receiver_id);rel(fk);null"`
	CreateTime time.Time `orm:"column(create_time);type(datetime);null"`
	Reward     float64   `orm:"column(reward)"`
	State      uint      `orm:"column(state);null"`
	Note       string    `orm:"column(note);size(200);null"`
}

func (t *Package) TableName() string {
	return "package"
}

func init() {
	orm.RegisterModel(new(Package))
}

func Typeof(v interface{}) string {
    return fmt.Sprintf("%T", v)
}

// AddPackage insert a new Package into database and returns
// last inserted Id on success.
func AddPackage(m *Package) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPackageById retrieves Package by Id. Returns error if
// Id doesn't exist
func GetPackageById(id int) (v *Package, err error) {
	o := orm.NewOrm()
	v = &Package{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPackage retrieves all Package matches certain condition. Returns empty list if
// no records exist
func GetAllPackage(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Package))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Package
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdatePackage updates Package by Id and returns error if
// the record to be updated doesn't exist
func UpdatePackageById(m *Package) (err error) {
	o := orm.NewOrm()
	v := Package{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePackage deletes Package by Id and returns error if
// the record to be deleted doesn't exist
func DeletePackage(id int) (err error) {
	o := orm.NewOrm()
	v := Package{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Package{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func ReceivePackage(id int, receiver_id int) (err error) {
	o := orm.NewOrm()
	v := Package{Id: id}
	if err = o.Read(&v); err == nil {
		v.State = 1
		v.ReceiverId, err = GetUserById(receiver_id)
		if err == nil {
			_, err = o.Update(&v, "State", "ReceiverId")
		}
	}
	return
}

func ConfirmPackage(id int) (err error) {
	o := orm.NewOrm()
	v := Package{Id: id}
	if err = o.Read(&v); err == nil {
		if v.OwnerId == nil || v.ReceiverId == nil {
			err = errors.New("owner or receiver doesn't exist")
			return
		}
		//感觉这里有一个orm自身的设计问题，ownerId除了id其他的竟然全为默认值
		v.OwnerId, _ = GetUserById(v.OwnerId.Id)
		v.ReceiverId, _ = GetUserById(v.ReceiverId.Id)
		//进行事务处理，避免资金转移错误
		o.Begin()
		//fmt.Println(v.OwnerId.Password)
		if v.OwnerId.Balance < v.Reward {
			err = errors.New(fmt.Sprintf("owner balance doesn't enough, balance:%f,reward:%f", v.OwnerId.Balance, v.Reward))
			return
		}
		v.OwnerId.Balance -= v.Reward
		v.ReceiverId.Balance += v.Reward
		err = UpdateUserById(v.OwnerId)
		if err != nil {
			o.Rollback()
			return
		}
		err = UpdateUserById(v.ReceiverId)
		if err != nil {
			o.Rollback()
			return
		}
		v.State = 2
		_, err = o.Update(&v, "state")
		if err != nil {
			o.Rollback()
			return
		}
		o.Commit()
	}
	return
}

func GetPackages(id int, owner_id int, receiver_id int, state int, limit int, offset int) []Package {
	var packages []Package
	o := orm.NewOrm()
	qs := o.QueryTable(new(Package))
	cond := orm.NewCondition()
	if id != 0 {
		cond = cond.And("Id", id)
	}
	if owner_id != 0 {
		user, err := GetUserById(owner_id)
		if err != nil {
			return packages
		}
		cond = cond.And("OwnerId", user)
	}
	if receiver_id != 0 {
		user, err := GetUserById(receiver_id)
		if err != nil {
			return packages
		}
		cond = cond.And("ReceiverId", user)
	}
	if state != -1 {
		cond = cond.And("State", state)
	}
	qs.SetCond(cond).Limit(limit).Offset(offset).All(&packages)
	return packages
}
