package controllers

import (
	"fmt"
	"github.com/make-money-sysu/server/models"
	"time"

	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson"
)

// PackageController operations for Package
type MsgController struct {
	beego.Controller
}

// 发送信息
func (this *MsgController) Post() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	this.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	
	var msg models.Msg
	bodyJSON := simplejson.New()
	if inputJson, err := simplejson.NewJson(this.Ctx.Input.RequestBody); err == nil {
		if nil == this.GetSession("id") {
			this.Ctx.Output.SetStatus(401)
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", "Login expired")
		}else{
			var err1 error
			var err2 error
			err1 = nil
			err2 = nil
			msg.Fromid,err1 = models.GetUserById(this.GetSession("id").(int))
			msg.Toid,err2 = models.GetUserById(inputJson.Get("to").MustInt())
			if err1 != nil || err2 != nil {// 查看用户是否存在
				// fmt.Println(inputJson.Get("to").MustInt())
				this.Ctx.Output.SetStatus(403)
				bodyJSON.Set("status", "failed")
				bodyJSON.Set("msg", "user not found")
			}else{
				msg.Content = inputJson.Get("msg").MustString()
				msg.Createtime = time.Now()
				msg.State = 10
				// 0为系统消息 ，10为未查看，11为已查看，但未知悉，12为已查看，已知悉，13为已撤回
			
				result, err := models.SendMessage(&msg)
				if err == nil {
					bodyJSON.Set("status", "success")
					bodyJSON.Set("msg", result)
				} else {
					this.Ctx.Output.SetStatus(403)
					bodyJSON.Set("status", "failed")
					bodyJSON.Set("msg", result)
				}
			}
		}
	} else {
		this.Ctx.Output.SetStatus(400)
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("msg", "invalid msg format")
	}
	body, _ := bodyJSON.Encode()
	this.Ctx.Output.Body(body)
}

// 撤回信息，（只能是未读的） WithdrawalMessage
func (this *MsgController) Delete() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	this.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	

	bodyJSON := simplejson.New()
	if inputJson, err := simplejson.NewJson(this.Ctx.Input.RequestBody); err == nil {
		if nil == this.GetSession("id") {
			this.Ctx.Output.SetStatus(401)
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", "Login expired")
		}else{
			fromid := this.GetSession("id").(int)
			mid := inputJson.Get("mid").MustInt()
			// 0为系统消息 ，10为未查看，11为已查看，但未知悉，12为已查看，已知悉，13为已撤回
		
			result, err := models.WithdrawalMessage(fromid,mid)
			if err == nil {
				bodyJSON.Set("status", "success")
				bodyJSON.Set("msg", result)
			} else {
				this.Ctx.Output.SetStatus(403)
				bodyJSON.Set("status", "failed")
				bodyJSON.Set("msg", result)
			}
		}
	} else {
		this.Ctx.Output.SetStatus(400)
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("msg", "invalid msg format")
	}
	body, _ := bodyJSON.Encode()
	this.Ctx.Output.Body(body)
}

//获取消息, 被获取了，数据库就算已读（TODO 状态修改还没加）
func (this *MsgController) Get() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	this.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	
	bodyJSON := simplejson.New()
	fmt.Println(this.GetSession("id"))
	if this.GetSession("id") == nil  {
		this.Ctx.Output.SetStatus(401)
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("msg", "Login expired")
	}else{
		var readData []models.Msg
		var unreadData []models.Msg
		fromid := this.GetSession("id").(int)
		// mode 0 is history, 1 (not zero) is for change
		if history, err :=this.GetInt("history"); err == nil {
			if history == 0 {

				limit, err := this.GetInt("limit")
				if err != nil {
					limit = -1
				}
				offset, err := this.GetInt("offset")
				if err != nil {
					offset = 0
				}
				toid,err := this.GetInt("with")
				if err == nil {
					readData, err = models.GetHistory(fromid,toid, limit, offset)
				}else{
					this.Ctx.Output.SetStatus(403)
					bodyJSON.Set("status", "failed")
					bodyJSON.Set("msg", "please give the people id whose record with you you want to get")
				}
			}else{
				readData,unreadData, err = models.GetMessage(fromid)
			}

			if err == nil {
				bodyJSON.Set("status", "success")
	
				tmpMapArr := make([]interface{}, len(readData))
				for i, p := range readData {
					tmpMap := make(map[string]interface{})
					tmpMap["mid"] = p.Mid
					tmpMap["fromid"] = p.Fromid.Id
					fromuser, _ := models.GetUserById(p.Fromid.Id)
					tmpMap["from_realname"] = fromuser.RealName
					tmpMap["from_nickname"] = fromuser.NickName
					tmpMap["toid"] = p.Toid.Id
					touser, _ := models.GetUserById(p.Toid.Id)
					tmpMap["to_realname"] = touser.RealName
					tmpMap["to_nickname"] = touser.NickName
					
					tmpMap["create_time"] = p.Createtime.String()
					tmpMap["Content"] = p.Content
					tmpMap["state"] = p.State
					tmpMapArr[i] = tmpMap
				}
				bodyJSON.Set("readData", tmpMapArr)
	
				if history != 0 {
					tmpMapArr := make([]interface{}, len(unreadData))
					for i, p := range unreadData {
						tmpMap := make(map[string]interface{})
						tmpMap["mid"] = p.Mid
						tmpMap["fromid"] = p.Fromid.Id
						fromuser, _ := models.GetUserById(p.Fromid.Id)
						tmpMap["from_realname"] = fromuser.RealName
						tmpMap["from_nickname"] = fromuser.NickName
						tmpMap["toid"] = p.Toid.Id
						touser, _ := models.GetUserById(p.Toid.Id)
						tmpMap["to_realname"] = touser.RealName
						tmpMap["to_nickname"] = touser.NickName

						tmpMap["create_time"] = p.Createtime.String()
						tmpMap["Content"] = p.Content
						tmpMap["state"] = p.State
						tmpMapArr[i] = tmpMap
					}
					bodyJSON.Set("unreadData", tmpMapArr)
				}
			} else {
				if _, ok := bodyJSON.CheckGet("status");!ok{
					this.Ctx.Output.SetStatus(403)
					bodyJSON.Set("status", "failed")
					bodyJSON.Set("msg", "get message error,you can try to find us for help")
				}
			}
		}else{
			this.Ctx.Output.SetStatus(400)
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", "please input history mode")
		}
	}
	
	body, _ := bodyJSON.Encode()
	this.Ctx.Output.Body(body)
}
