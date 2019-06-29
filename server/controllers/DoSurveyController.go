package controllers

import (
	"github.com/make-money-sysu/server/models"
	"time"

	// "fmt"

	"github.com/bitly/go-simplejson"

	"github.com/astaxie/beego"
)

type DoSurveyController struct {
	beego.Controller
}

// @Title Get
// @Description 用来获取填写问卷的信息列表
// @Param	survey_id							query		int		false	"问卷id"
// @Param	recipient_id						query		int		false	"填写问卷者id"
// @Param	content								query		string	false	"填写问卷的内容"
// @Param	create_time							query		string	false	"填写问卷的时间"
// @Success	200				{"status" : "success", "data": {填写问卷的信息列表数组}}
// @router / [get]
func (this *DoSurveyController) Get() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	this.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	surver_id, err := this.GetInt("survey_id")
	if err != nil {
		surver_id = -1
	}
	recipient_id, err := this.GetInt("recipient_id")
	if err != nil {
		recipient_id = -1
	}
	content := this.GetString("content")
	create_time := this.GetString("create_time")
	records := models.QueryDoSurvey(surver_id, recipient_id, content, create_time)
	bodyJSON := simplejson.New()
	bodyJSON.Set("status", "success")
	tmpMapArr := make([]interface{}, len(records))
	for i, r := range records {
		tmpMap := make(map[string]interface{})
		tmpMap["survey_id"] = r.SurveyId.Id
		tmpMap["recipient_id"] = r.RecipientId.Id
		tmpMap["content"] = r.Content
		tmpMap["create_time"] = r.CreateTime.Format("2006-01-02 15:04:05")
		tmpMapArr[i] = tmpMap
	}
	bodyJSON.Set("data", tmpMapArr)
	body, _ := bodyJSON.Encode()
	this.Ctx.Output.Body(body)
}

// @Title Post
// @Description 用来上传填写问卷的信息
// @Param	survey_id							body		int		true	"问卷id"
// @Param	recipient_id						body		int		true	"填写问卷者id"
// @Param	content								body		string	true	"填写问卷的内容"
// @Success	200				{"status" : "success", "msg": "add do survey record succeed"}
// @Failure 400				{"status" : "failed", "msg": "invalid survey id"}
// @Failure 400				{"status" : "failed", "msg": "invalid user id"}
// @Failure 400				{"status" : "failed", "msg": "invalid do survey json format"}
// @router / [post]
func (this *DoSurveyController) Post() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	this.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	if inputJSON, err := simplejson.NewJson(this.Ctx.Input.RequestBody); err == nil {
		survey_id := inputJSON.Get("survey_id").MustInt()
		recipient_id := inputJSON.Get("recipient_id").MustInt()
		content := inputJSON.Get("content").MustString()
		//loc, _ := time.LoadLocation("UTC")
		Time := time.Now().Local()
		//fmt.Println(Time.Format("2006-01-02 15:04:05"))
		var do_survey models.DoSurvey
		bodyJSON := simplejson.New()
		do_survey.SurveyId, err = models.GetSurveyById(survey_id)
		if err != nil {
			this.Ctx.Output.SetStatus(400)
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", "invalid survey id")
			body, _ := bodyJSON.Encode()
			this.Ctx.Output.Body(body)
			return
		}
		do_survey.RecipientId, err = models.GetUserById(recipient_id)
		if err != nil {
			this.Ctx.Output.SetStatus(400)
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", "invalid user id")
			body, _ := bodyJSON.Encode()
			this.Ctx.Output.Body(body)
			return
		}
		do_survey.Content = content
		do_survey.CreateTime = Time
		err := models.CreateDoSurvey(&do_survey)
		if err != nil {
			this.Ctx.Output.SetStatus(400)
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", err.Error())
			body, _ := bodyJSON.Encode()
			this.Ctx.Output.Body(body)
			return
		} else {
			bodyJSON.Set("status", "success")
			bodyJSON.Set("msg", "add do survey record succeed")
			body, _ := bodyJSON.Encode()
			this.Ctx.Output.Body(body)
		}
	} else {
		bodyJSON := simplejson.New()
		this.Ctx.Output.SetStatus(400)
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("msg", "invalid do survey json format")
		body, _ := bodyJSON.Encode()
		this.Ctx.Output.Body(body)
	}
}
