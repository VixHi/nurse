package router

import (
	"fmt"
	"viv/controllers"

	"github.com/astaxie/beego"
)

func init() {
	fmt.Println("=======router==========")

	//nurse 路由
	nurseNS := beego.NewNamespace("/nurse",

		beego.NSRouter("", &controllers.NurseController{}),

		beego.NSRouter("/title", &controllers.NurseController{}, "get:GetNurseTitle;Put:UpdateNurseTitle"),
	)
	beego.AddNamespace(nurseNS)

	loginNS := beego.NewNamespace("/login",
		beego.NSRouter("/verifyCode", &controllers.LoginController{}, "get:GetVerifyCode;post:CreateVerifyCode"),
		beego.NSRouter("/password", &controllers.LoginController{}, "post:GetPwd"),
		beego.NSRouter("/images", &controllers.LoginController{}, "post:UploageUserImage"),
		beego.NSRouter("/users", &controllers.LoginController{}, "post:RegisterUser"),
	)
	beego.AddNamespace(loginNS)
}
