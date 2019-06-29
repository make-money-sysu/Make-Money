package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type DoSurvey struct {
	Id          int       `orm:"column(id);auto"`
	SurveyId    *Survey   `orm:"column(survey_id);rel(fk)"`
	RecipientId *User     `orm:"column(recipient_id);rel(fk)"`
	Content     string    `orm:"column(content);size(1000)"`
	CreateTime  time.Time `orm:"column(create_time);type(datetime);null"`
}

func (t *DoSurvey) TableName() string {
	return "do_survey"
}

func init() {
	orm.RegisterModel(new(DoSurvey))
}

// AddDoSurvey insert a new DoSurvey into database and returns
// last inserted Id on success.
func AddDoSurvey(m *DoSurvey) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDoSurveyById retrieves DoSurvey by Id. Returns error if
// Id doesn't exist
func GetDoSurveyById(id int) (v *DoSurvey, err error) {
	o := orm.NewOrm()
	v = &DoSurvey{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDoSurvey retrieves all DoSurvey matches certain condition. Returns empty list if
// no records exist
func GetAllDoSurvey(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(DoSurvey))
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

	var l []DoSurvey
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

// UpdateDoSurvey updates DoSurvey by Id and returns error if
// the record to be updated doesn't exist
func UpdateDoSurveyById(m *DoSurvey) (err error) {
	o := orm.NewOrm()
	v := DoSurvey{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDoSurvey deletes DoSurvey by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDoSurvey(id int) (err error) {
	o := orm.NewOrm()
	v := DoSurvey{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&DoSurvey{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func QueryDoSurvey(survey_id int, recipient_id int, content string, create_time string) []DoSurvey {
	fmt.Println(create_time)
	var records []DoSurvey
	o := orm.NewOrm()
	cond := orm.NewCondition()
	if survey_id != -1 {
		survey, err := GetSurveyById(survey_id)
		if err != nil {
			return records
		}
		cond = cond.And("SurveyId", survey)
	}
	if recipient_id != -1 {
		user, err := GetUserById(recipient_id)
		if err != nil {
			return records
		}
		cond = cond.And("RecipientId", user)
	}
	loc, _ := time.LoadLocation("Local")
	time, err := time.ParseInLocation("2006-01-02 15:04:05", create_time, loc)
	if err == nil {
		cond = cond.And("CreateTime", time)
	}
	qs := o.QueryTable(new(DoSurvey))
	qs.SetCond(cond).Filter("Content__contains", content).All(&records)
	return records

}

func CreateDoSurvey(d *DoSurvey) error {
	survey, _ := GetSurveyById(d.SurveyId.Id)
	recipient, _ := GetUserById(d.RecipientId.Id)
	owner, _ := GetUserById(survey.PublisherId.Id)
	o := orm.NewOrm()
	o.Begin()
	if owner.Balance < 0.1 {
		return errors.New("owner balance not enough")
	}
	owner.Balance -= 0.1
	recipient.Balance += 0.1
	err := UpdateUserById(owner)
	if err != nil {
		o.Rollback()
		return err
	}
	err = UpdateUserById(recipient)
	if err != nil {
		o.Rollback()
		return err
	}
	_, err = AddDoSurvey(d)
	if err != nil {
		o.Rollback()
		return err
	}
	o.Commit()
	return nil

}
