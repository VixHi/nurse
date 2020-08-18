package controllers

import (
	"strconv"
	"viv/models"
	"viv/vutil"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// NewsController : 新闻控制器
type NewsController struct {
	beego.Controller
}

// Get : Get请求
func (c *NewsController) Get() {
	//多对多写入
	newsID := c.GetString("newsId")
	ID, err := strconv.Atoi(newsID)
	if err != nil {
		beego.Info(err)
		return
	}
	news := models.News{
		Id: ID,
	}
	o := orm.NewOrm()
	/* //多对多写入
	m2m := o.QueryM2M(&news, "Nurses")
	nurse := models.Nurse{
		Id: 8,
	}
	o.Read(&nurse)
	_, err = m2m.Add(&nurse)
	if err != nil {
		beego.Info(err)
		return
	}
	*/

	//多对多查询
	err = o.QueryTable("news").RelatedSel().One(&news)
	if err != nil {
		beego.Info(err)
		return
	}
	beego.Info(news)

	o.LoadRelated(&news, "Nurses")
	c.Data["json"] = vutil.ResponseWith(200, "success", news)
	c.ServeJSON()
	// c.Ctx.WriteString("成功")
}

// Post : 新闻发布
func (c *NewsController) Post() {
	news := models.News{
		Title:   "新闻标题2",
		Content: "新闻内容2",
	}

	hospitalID := c.GetString("hospitalId")
	ID, err := strconv.Atoi(hospitalID)
	hospital := models.Hospital{
		Id: ID,
	}

	o := orm.NewOrm()
	err = o.Read(&hospital)
	if err != nil {
		beego.Info(err)
		return
	}
	news.Hospital = &hospital
	beego.Info(news)
	_, err = o.Insert(&news)
	if err != nil {
		beego.Info(err)
		return
	}
	c.Data["json"] = vutil.ResponseWith(200, "success", news)
	c.ServeJSON()
}
