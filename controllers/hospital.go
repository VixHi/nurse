package controllers

import (
	"strconv"
	"viv/models"
	"viv/vutil"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// HospitalController : 医生控制器
type HospitalController struct {
	beego.Controller
}

// Get : 医院获取
func (c *HospitalController) Get() {
	nurseID := c.GetString("nurseId")
	if nurseID != "" {
		getHospitalByNurse(c, nurseID)
		return
	}
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

// getHospitalByNurse : 根据护士检索医院
func getHospitalByNurse(c *HospitalController, nurseID string) {
	ID, err := strconv.Atoi(nurseID)
	if err != nil {
		beego.Info(err)
		return
	}
	nurse := models.Nurse{
		Id: ID,
	}
	o := orm.NewOrm()

	err = o.QueryTable("nurse").Filter("Id", nurseID).RelatedSel().One(&nurse)
	if err != nil {
		beego.Info(err)
		return
	}
	_, err = o.LoadRelated(&nurse, "Hospital")
	hospital := nurse.Hospital
	c.Data["json"] = vutil.ResponseWith(200, "success", hospital)
	c.ServeJSON()
}

// getHospitalByID : 根据医院ID检索医院
func getHospitalByID() {

}
