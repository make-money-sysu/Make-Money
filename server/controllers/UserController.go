package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/make-money-sysu/server/models"
	// "strconv"

	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson"
)

type UserController struct {
	beego.Controller
}

// @Title Post
// @Description 用来注册用户，body为json格式
// @Param	id			body		int		true	"用户的学号"
// @Param	password	body		string	true	"用户的密码"
// @Param 	real_name	body		string	true	"用户的真名"
// @Param	nick_name	body		string	true	"用户的昵称"
// @Param	age			body		uint16	true	"用户的年龄"
// @Param	gender		body		string	true	"用户的性别"
// @Param	head_piture	body		string	true	"用户头像（从字节转换成字符串后）"
// @Param	balance		body		float32	true	"用户存款"
// @Param	profession	body		string	true	"用户专业"
// @Param	grade		body		string	true	"用户年级"
// @Param	phone		body		string	true	"用户电话"
// @Param	email		body		string	true	"用户邮箱"
// @Success	200			{"status" : "success", "msg": "resgiter succeed"}
// @Failure	403			{"status" : "failed", "msg": "this user already registered"}
// @Failure 400			{"status" : "failed", "msg": "invalid user infomation format", "err" : {some error}}
// @router / [post]
func (this *UserController) Post() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	this.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")

	var user models.User
	bodyJSON := simplejson.New()
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &user); err == nil {
		// fmt.Printf("add user: %+v\n", user)
		_, err = models.AddUser(&user)
		if err == nil {
			bodyJSON.Set("status", "success")
			bodyJSON.Set("msg", "resgiter succeed")
			// this.SetSession("id", id)
		} else {
			this.Ctx.Output.SetStatus(403)
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", "this user already registered")
		}
	} else {
		this.Ctx.Output.SetStatus(400)
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("msg", "invalid user infomation format")
		bodyJSON.Set("err", err)
		fmt.Println(err)
	}
	body, _ := bodyJSON.MarshalJSON()
	this.Ctx.Output.Body(body)
}

// @Title Put
// @Description 用来修改用户信息，body为json格式
// @Param	id			body		int		true	"用户的学号"
// @Param	password	body		string	true	"用户的密码"
// @Param 	real_name	body		string	true	"用户的真名"
// @Param	nick_name	body		string	true	"用户的昵称"
// @Param	age			body		uint16	true	"用户的年龄"
// @Param	gender		body		string	true	"用户的性别"
// @Param	head_piture	body		string	true	"用户头像（从字节转换成字符串后）"
// @Param	balance		body		float32	true	"用户存款"
// @Param	profession	body		string	true	"用户专业"
// @Param	grade		body		string	true	"用户年级"
// @Param	phone		body		string	true	"用户电话"
// @Param	email		body		string	true	"用户邮箱"
// @Success	200			{"status" : "success", "msg": "edited"}
// @Failure 401			{"status" : "failed", "msg": "Login expired"}
// @Failure 403			{"status" : "failed", "msg": "this user doesn't exist"}
// @Failure 400			{"status" : "failed", "msg": "invalid user infomation format"}
// @router / [put]
func (this *UserController) Put() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	this.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	
	fmt.Println(this.Ctx.Input.Header("cookie"))

	bodyJSON := simplejson.New()
	if inputJSON, err := simplejson.NewJson(this.Ctx.Input.RequestBody); err == nil {
		if nil == this.GetSession("id") {
			this.Ctx.Output.SetStatus(401)
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", "Login expired")
		}else{
			user, err := models.GetUserById(this.GetSession("id").(int))

			fmt.Println(user)
			fmt.Println(inputJSON)

			if _, ok := inputJSON.CheckGet("real_name");ok{
				user.RealName = inputJSON.Get("real_name").MustString()
			}
	
			if _, ok := inputJSON.CheckGet("password");ok{
				user.Password = inputJSON.Get("password").MustString()
			}
	
			if _, ok := inputJSON.CheckGet("nick_name");ok{
				user.NickName = inputJSON.Get("nick_name").MustString()
			}
	
			if _, ok := inputJSON.CheckGet("age");ok{
				user.Age = inputJSON.Get("age").MustInt()
			}
	
			if _, ok := inputJSON.CheckGet("gender");ok{
				user.Gender = inputJSON.Get("gender").MustString()
			}
	
			if _, ok := inputJSON.CheckGet("head_picture");ok{
				user.HeadPicture = inputJSON.Get("head_picture").MustString()
			}
	
			// if _, ok := inputJSON.CheckGet("balance");ok{
			// 	user.Balance = inputJSON.Get("balance").MustFloat64()
			// }
	
	
			if _, ok := inputJSON.CheckGet("profession");ok{
				user.Profession = inputJSON.Get("profession").MustString()
			}
	
			
			if _, ok := inputJSON.CheckGet("grade");ok{
				user.Grade = inputJSON.Get("grade").MustString()
			}
	
			if _, ok := inputJSON.CheckGet("phone");ok{
				user.Phone = inputJSON.Get("phone").MustString()
			}
	
			if _, ok := inputJSON.CheckGet("email");ok{
				user.Email = inputJSON.Get("email").MustString()
			}
	
		
			
			fmt.Println(user)

			err = models.UpdateUserById(user)
			if err == nil {
				bodyJSON.Set("status", "success")
				bodyJSON.Set("msg", "edited")
			} else {
				this.Ctx.Output.SetStatus(403)
				bodyJSON.Set("status", "failed")
				bodyJSON.Set("msg", "this user doesn't exist")
			}
		}
	} else {
		this.Ctx.Output.SetStatus(400)
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("msg", "invalid user infomation format")
	}
	body, _ := bodyJSON.MarshalJSON()
	this.Ctx.Output.Body(body)
}

// @Title Delete
// @Description 用来删除当前登录的用户
// @Success	200			{"status" : "success", "msg": "bye~"}
// @Failure 401			{"status" : "failed", "msg": "Login expired"}
// @Failure 403			{"status" : "failed", "msg": "invalid user"}
// @router / [delete]
func (this *UserController) Delete() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	this.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")


	bodyJSON := simplejson.New()

	if nil == this.GetSession("id") {
		this.Ctx.Output.SetStatus(401)
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("msg", "Login expired")
	} else {
		id := this.GetSession("id").(int)
		err := models.DeleteUser(id)

		if err == nil {
			bodyJSON.Set("status", "success")
			bodyJSON.Set("msg", "bye~")
		} else {
			this.Ctx.Output.SetStatus(403)
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", "invalid user")
		}
	}

	body, _ := bodyJSON.Encode()
	this.Ctx.Output.Body(body)
}

// @Title Get
// @Description 用来获取当前登录用户的用户信息
// @Success	200			{"status" : "success", "data": {json格式的用户信息}}
// @Failure 400			{"status" : "failed", "msg": "Login expired"}
// @Failure 401			{"status" : "failed", "msg": "user doesn't exist"}
// @router / [get]
func (this *UserController) Get() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	this.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")


	bodyJSON := simplejson.New()
	//id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
	fmt.Println(this.GetSession("id"))
	if this.GetSession("id") != nil {
		id := this.GetSession("id").(int)
		var user *models.User
		user, err := models.GetUserById(id)
		if err == nil {
			bodyJSON.Set("status", "success")
			dataMap := make(map[string]interface{})
			dataMap["id"] = id
			dataMap["real_name"] = user.RealName
			dataMap["nick_name"] = user.NickName
			dataMap["age"] = user.Age
			dataMap["gender"] = user.Gender
			dataMap["head_picture"] = user.HeadPicture
			dataMap["balance"] = user.Balance
			dataMap["profession"] = user.Profession
			dataMap["grade"] = user.Grade
			dataMap["phone"] = user.Phone
			dataMap["email"] = user.Email
			bodyJSON.Set("data", dataMap)
		} else {
			this.Ctx.Output.SetStatus(401)
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", "user doesn't exist")
		}
	} else {
		this.Ctx.Output.SetStatus(400)
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("msg", "Login expired")
	}
	body, _ := bodyJSON.MarshalJSON()
	this.Ctx.Output.Body(body)
}
