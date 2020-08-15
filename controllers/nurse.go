package controllers

import "github.com/astaxie/beego"

// NurseController : 护士控制器
type NurseController struct {
	beego.Controller
}

// Get : 获取用户信息
func (c *NurseController) Get() {

	c.Ctx.WriteString("1234")
}

// Post : 新增护士
func (c *NurseController) Post() {
	c.Ctx.WriteString("1234")

}

// Put : 更改用户信息
func (c *NurseController) Put() {
	c.Ctx.WriteString("1234")

}

// Delete : 删除护士
func (c *NurseController) Delete() {
	c.Ctx.WriteString("12345")

}

// GetNurseTitle : 获取护士的职称
func (c *NurseController) GetNurseTitle() {
	c.Ctx.WriteString("获取护士的职称")
}

// UpdateNurseTitle :更改护士的职称
func (c *NurseController) UpdateNurseTitle() {
	c.Ctx.WriteString("更改护士的职称")
}
