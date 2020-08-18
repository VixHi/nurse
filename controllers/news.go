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
