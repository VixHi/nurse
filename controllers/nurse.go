package controllers

import (
	"strconv"
	"viv/models"
	"viv/vutil"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// NurseController : 护士控制器
type NurseController struct {
	beego.Controller
}

// nurseParse : 参数解析到结构体，通过form形式
type nurseParse struct {
	Page       int    `form:"page"`
	Size       int    `form:"size"`
	Title      string `form:"title"`
	HospitalId string `form:"hospitalId"`
	NurseId    string `form:"nurseId"`
}

//全局变量
var nurse_parse nurseParse //护士请求参数

// Get : 获取用户信息
func (c *NurseController) Get() {
	nurse_parse = nurseParse{}
	if err := c.ParseForm(&nurse_parse); err != nil {
		beego.Info(err)
		return
	}
	beego.Info(nurse_parse)

	nurseID := c.GetString("id")
	if nurseID != "" {
		getNurseByID(c, nurseID)
		return
	}
	nurseTitle := c.GetString("title")
	if nurseTitle != "" {
		getNursesByTitle(c, nurseTitle)
		return
	}
	hospitalID := c.GetString("hospitalId")
	if hospitalID != "" {
		getNursesByHospitalID(c, hospitalID)
		return
	}
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

// GetNursesByType : 获取当前类型的全部护士
func (c *NurseController) GetNursesByType() {

}

// getNurseByID : 通过id获取护士
func getNurseByID(c *NurseController, id string) {
	nurseID, err := strconv.Atoi(id)
	if err != nil {
		beego.Info(err)
		return
	}
	nurse := models.Nurse{
		Id: nurseID,
	}

	o := orm.NewOrm()
	err = o.Read(&nurse)
	if err != nil {
		beego.Info(err)
		return
	}
	beego.Info(nurse)
	c.Data["json"] = vutil.ResponseWith(200, "数据请求成功", nurse)
	c.ServeJSON()

}

// getNurses : 获取全部title下的护士
func getNursesByTitle(c *NurseController, nurseTitle string) {

	nurses := []models.Nurse{}
	o := orm.NewOrm()
	_, err := o.QueryTable("nurse").Filter("title", nurseTitle).All(&nurses)
	if err != nil {
		beego.Info(err)
		return
	}
	beego.Info(nurses)
	c.Data["json"] = vutil.ResponseWith(200, "success", nurses)
	c.ServeJSON()
}

// getNursesByHospitalID : 获取医院全部护士
func getNursesByHospitalID(c *NurseController, hospitalID string) {
	ID, err := strconv.Atoi(hospitalID)
	if err != nil {
		beego.Info(err)
		return
	}
	hospital := models.Hospital{
		Id: ID,
	}
	nurses := []models.Nurse{}
	o := orm.NewOrm()
	//外健 获取hospial数据
	qs := o.QueryTable("nurse")
	total, err := qs.Filter("Hospital", hospital).RelatedSel().Count()
	if err != nil {
		beego.Info(err)
		return
	}
	// _, err = o.RelatedSel(&nurses, "Hospital")
	beego.Info("护士总数:", total)
	qs.Limit(nurse_parse.Size, 1).All(&nurses)
	c.Data["json"] = vutil.ResponseWith(200, "success", nurses)
	c.ServeJSON()
}
