// @APIVersion 1.0.0
// @Title make-money api
// @Description a api for make-money web application
// @Contact 935841375@qq.com
package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/make-money-sysu/server/controllers"
)

func Init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	ns := beego.NewNamespace("/api",

		beego.NSNamespace("/do_survey",
			beego.NSInclude(
				&controllers.DoSurveyController{},
			),
		),

		beego.NSNamespace("/friends",
			beego.NSInclude(
				&controllers.FriendsController{},
			),
		),

		beego.NSNamespace("/msg",
			beego.NSInclude(
				&controllers.MsgController{},
			),
		),

		beego.NSNamespace("/package",
			beego.NSInclude(
				&controllers.PackageController{},
			),
		),

		beego.NSNamespace("/survey",
			beego.NSInclude(
				&controllers.SurveyController{},
			),
		),

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),

		beego.NSNamespace("/login",
			beego.NSInclude(
				&controllers.LoginController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
