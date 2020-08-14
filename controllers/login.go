package controllers

import (
	"path"
	"strconv"
	"time"
	"viv/models"
	"viv/variable"
	"viv/vutil"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/gomodule/redigo/redis"
)

// LoginController : 登录控制器
type LoginController struct {
	beego.Controller
}

// GetVerifyCode : 获取验证码
func (c *LoginController) GetVerifyCode() {
	c.Ctx.WriteString("获取验证码")
	//获取随机验证码
	/*
		cac, err := cache.NewCache("redis", `{"conn":"127.0.0.1:6379", "key":"beecacheRedis"}`)
		if err != nil {
			beego.Info(err)
			return
		}
		code := cac.Get("verifyCode")
		beego.Info(code)
		c.Data["json"] = map[string]interface{}{"code": code}
		c.ServeJSON()
	*/

	// beego.Info("获取参数code:", c.Ctx.Input.Param("code"), c.Input().Get("code"), c.Ctx.Request)
	//获取用户传入的code
	userCode := c.Input().Get("code")

	//服务器code
	conn, err := redis.Dial("tcp", ":6379")
	defer conn.Close()
	if err != nil {
		beego.Info(err)
		return
	}

	reply, err := redis.String(conn.Do("get", "code"))
	if err != nil {
		beego.Info(err)
		return
	}

	if reply == userCode {
		c.Data["json"] = map[string]string{
			"msg":   "成功",
			"code":  "200",
			"error": "",
		}
	} else {
		c.Data["json"] = map[string]string{
			"msg":   "失败",
			"code":  "400",
			"error": variable.VphoneOrVerifyCodeErr,
		}
	}

	c.ServeJSON()
}

// CreateVerifyCode : 生成验证码
func (c *LoginController) CreateVerifyCode() {
	//生成随机数 模仿验证码
	// rand.Seed(time.Now().UnixNano())
	// code := rand.Intn(8999) + 1000
	code := vutil.RandNum(1000, 9999)

	conn, err := redis.Dial("tcp", ":6379")
	defer conn.Close()
	if err != nil {
		beego.Info(err)
		return
	}

	conn.Do("set", "code", code)
	phone := c.GetString("phone")
	beego.Info(phone)
	conn.Do("set", "phone", phone)
	codeStr := strconv.Itoa(code)
	c.Data["json"] = map[string]string{
		"phone": phone,
		"code":  codeStr,
	}
	c.ServeJSON()

	c.Ctx.WriteString(strconv.Itoa(code))

	/*
		beego.Info(code)
		cac, err := cache.NewCache("redis", `{"conn":"127.0.0.1:6379", "key":"beecacheRedis"}`)
		if err != nil {
			beego.Info(err)
			return
		}
		cac.Put("verifyCode", code, 1800*time.Second)
	*/

}

// GetPwd : 获取注册密码并保存
func (c *LoginController) GetPwd() {
	pwd := c.Input().Get("pwd")
	_ = vutil.CheckPwd(pwd)
	conn, err := redis.Dial("tcp", ":6379")
	defer conn.Close()
	if err != nil {
		beego.Info(err)
		return
	}
	//设置过期时间
	conn.Do("setex", "pwd", pwd, 1800*time.Second)
	c.Data["json"] = map[string]string{
		"pwd": pwd,
	}
	c.ServeJSON()
}

// UploageUserImage : 护士资格证上传
func (c *LoginController) UploageUserImage() {

	f, h, err := c.GetFile("userImage")
	if err != nil {
		beego.Info(err)
		return
	}
	defer f.Close()
	fileExt := path.Ext(h.Filename)
	if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".jpeg" {
		beego.Info("文件格式不对")
		return
	}

	if h.Size > 5*1024*1024 {
		beego.Info("文件太大")
		return
	}

	timeStr := time.Now().Format("2006-01-02 03:04:06")
	imageName := timeStr + fileExt
	imageURL := "/static/img/" + imageName
	beego.Info(imageURL)
	c.SaveToFile("userImage", "."+imageURL)
	c.Ctx.WriteString("文件上传成功")

	o := orm.NewOrm()
	img := models.Image{
		ImageStr: imageURL,
	}
	_, err = o.Insert(&img)
	if err != nil {
		beego.Info(err)
		return
	}
}

// RegisterUser : 用户注册
func (c *LoginController) RegisterUser() {
	conn, err := redis.Dial("tcp", ":6379")
	defer conn.Close()
	if err != nil {
		beego.Info(err)
		c.Data["json"] = vutil.ResponseWith(500, "数据库连接错误", nil)
		return
	}

	phone, err := redis.String(conn.Do("get", "phone"))
	beego.Info(phone)
	if err != nil {
		beego.Info(err)
		c.Data["json"] = vutil.ResponseWith(500, "数据库连接错误", nil)
		return
	}

	nurse := models.Nurse{
		Name:  "vix" + phone,
		Phone: phone,
		Sex:   "女",
		Title: "主任护士",
	}

	o := orm.NewOrm()
	_, err = o.Insert(&nurse)
	if err != nil {
		beego.Info(err)
		c.Data["json"] = vutil.ResponseWith(500, "数据库插入出错", nil)
		return
	}

	c.Data["json"] = vutil.ResponseWith(500, "数据库插入出错", nurse)

}
