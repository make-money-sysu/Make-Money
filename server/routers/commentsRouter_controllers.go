package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/make-money-sysu/server/controllers:DoSurveyController"] = append(beego.GlobalControllerRouter["github.com/make-money-sysu/server/controllers:DoSurveyController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/make-money-sysu/server/controllers:DoSurveyController"] = append(beego.GlobalControllerRouter["github.com/make-money-sysu/server/controllers:DoSurveyController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/make-money-sysu/server/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/make-money-sysu/server/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/make-money-sysu/server/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/make-money-sysu/server/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/make-money-sysu/server/controllers:PackageController"] = append(beego.GlobalControllerRouter["github.com/make-money-sysu/server/controllers:PackageController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/make-money-sysu/server/controllers:PackageController"] = append(beego.GlobalControllerRouter["github.com/make-money-sysu/server/controllers:PackageController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/make-money-sysu/server/controllers:PackageController"] = append(beego.GlobalControllerRouter["github.com/make-money-sysu/server/controllers:PackageController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/make-money-sysu/server/controllers:SurveyController"] = append(beego.GlobalControllerRouter["github.com/make-money-sysu/server/controllers:SurveyController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/make-money-sysu/server/controllers:SurveyController"] = append(beego.GlobalControllerRouter["github.com/make-money-sysu/server/controllers:SurveyController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/make-money-sysu/server/controllers:SurveyController"] = append(beego.GlobalControllerRouter["github.com/make-money-sysu/server/controllers:SurveyController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/make-money-sysu/server/controllers:SurveyController"] = append(beego.GlobalControllerRouter["github.com/make-money-sysu/server/controllers:SurveyController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/make-money-sysu/server/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/make-money-sysu/server/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/make-money-sysu/server/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/make-money-sysu/server/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/make-money-sysu/server/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/make-money-sysu/server/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/make-money-sysu/server/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/make-money-sysu/server/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
