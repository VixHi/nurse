package controllers

import "github.com/astaxie/beego"

// NurseController : 护士控制器
type NurseController struct {
	beego.Controller
}

// Get : 获取用户信息
func (nurse *NurseController) Get() {
	nurse.Ctx.WriteString("1234")
}

// Post : 新增护士
func (nurse *NurseController) Post() {
	nurse.Ctx.WriteString("1234")

}

// Put : 更改用户信息
func (nurse *NurseController) Put() {
	nurse.Ctx.WriteString("1234")

}

// Delete : 删除护士
func (nurse *NurseController) Delete() {
	nurse.Ctx.WriteString("12345")

}
