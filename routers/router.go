package router

import (
	"viv/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//nurse 路由
	beego.Router("/nurse", &controllers.NurseController{})
	//护士title路由
	// ns_nurse := beego.NewNameSpace("nurse",
	// 	beego.NSNameSpace("/title",
	// 		beego.NSIncloud(
	// 			&controllers.NurseController{}
	// 		),
	// 	),
	// )
	// beego.AddNameSpace(ns_nurse)
}
