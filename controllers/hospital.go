package controllers

import (
	"viv/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// HospitalController : 医生控制器
type HospitalController struct {
	beego.Controller
}

// Get : 医院获取
func (c *HospitalController) Get() {

}

// Post : 新增医院
func (c *HospitalController) Post() {

}

// Put : 更新医院
func (c *HospitalController) Put() {
	hospital := models.Hospital{Id: 4}
	o := orm.NewOrm()
	err := o.Read(&hospital)
	if err != nil {
		beego.Info(err)
		return
	}
	hospital.HospitalName = "第一中心医院"
	_, err = o.Update(&hospital, "HospitalName")
	if err != nil {
		beego.Info(err)
		return
	}
	c.Data["json"] = map[string]string{
		"code":    "200",
		"message": "更新成功",
	}
	c.ServeJSON()

}

// Delete : 删除医院
func (c *HospitalController) Delete() {
	hospital := models.Hospital{Id: 3}
	o := orm.NewOrm()
	_, err := o.Delete(&hospital)
	if err != nil {
		beego.Info(err)
		return
	}
	c.Ctx.WriteString("删除成功")
}
