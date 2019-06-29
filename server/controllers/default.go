package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
  c.AllowCross()
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

// Cross Domain
func (c *MainController) AllowCross() {
  c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin:*", "http://localhost:8080")       //允许访问源
  c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")    //允许post访问
  c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization") //header的类型
  c.Ctx.ResponseWriter.Header().Set("Access-Control-Max-Age", "1728000")
  c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
  c.Ctx.ResponseWriter.Header().Set("content-type", "application/json") //返回数据格式是json

}
