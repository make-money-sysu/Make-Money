package controllers

import (
	"fmt"
	"github.com/make-money-sysu/Make-Money/server/models"

	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson"
)

type LoginController struct {
	beego.Controller
}

// @Title Post
// @Description 用来登录
// @Param	id							body		int		true	"用户学号"
// @Param	password					body		string	true	"用户密码"
// @Success	200				{"status" : "success", "msg": "post success"}
// @Failure 400				{"status" : "failed", "msg": "invalid login format"}
// @Failure 403				{"status" : "failed", "msg": "id and password doesn't match"}
// @router / [post]
func (this *LoginController) Post() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	this.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")

	bodyJSON := simplejson.New()
	fmt.Println(this.Ctx.Input.Header("cookie"))
	if inputJSON, err := simplejson.NewJson(this.Ctx.Input.RequestBody); err == nil {
		id, err1 := inputJSON.Get("id").Int()
		passord, err2 := inputJSON.Get("password").String()
		if err1 != nil || err2 != nil {
			this.Ctx.Output.SetStatus(400)
			bodyJSON.Set("status", "fail")
			bodyJSON.Set("msg", "invalid login format")
		} else {
			success := models.UserLogin(id, passord)
			if success {
				bodyJSON.Set("status", "success")
				bodyJSON.Set("msg", "success")
				this.SetSession("id", id)
			} else {
				this.Ctx.Output.SetStatus(403)
				bodyJSON.Set("status", "failed")
				bodyJSON.Set("msg", "id and password doesn't match")
			}
		}
	} else {
		this.Ctx.Output.SetStatus(400)
		bodyJSON.Set("status", "fail")
		bodyJSON.Set("msg", "invalid login format")
	}
	body, _ := bodyJSON.MarshalJSON()
	this.Ctx.Output.Body(body)
}

// @Title Delete
// @Description 用来注销
// @Success	200				{"status" : "success", "msg": "logout succeed"}
// @router / [delete]
func (this *LoginController) Delete() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	this.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	this.DelSession("id")
	bodyJSON := simplejson.New()
	bodyJSON.Set("status", "success")
	bodyJSON.Set("msg", "logout succeed")
	body, _ := bodyJSON.MarshalJSON()
	this.Ctx.Output.Body(body)
}
