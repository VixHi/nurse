package router

import (
	"viv/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//nurse 路由
	nurseNS := beego.NewNamespace("/nurse",

		beego.NSRouter("", &controllers.NurseController{}),

		beego.NSRouter("/title", &controllers.NurseController{}, "get:GetNurseTitle;Put:UpdateNurseTitle"),

		// beego.NSNamespace("/api",
		// 	beego.NSInclude(
		// 		&controllers.NurseController{},
		// 	),
		// ),
	)
	beego.AddNamespace(nurseNS)

	loginNS := beego.NewNamespace("/login",
		beego.NSRouter("/verifyCode", &controllers.LoginController{}, "get:GetVerifyCode;post:CreateVerifyCode"),
		beego.NSRouter("/password", &controllers.LoginController{}, "post:GetPwd"),
		beego.NSRouter("/userImage", &controllers.LoginController{}, "post:UploageUserImage"),
	)
	beego.AddNamespace(loginNS)
}
