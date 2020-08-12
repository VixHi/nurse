package controllers

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/astaxie/beego"
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

	reply, err := redis.String(conn.Do("get", "codeVerify"))
	if err != nil {
		beego.Info(err)
		return
	}

	fmt.Printf("%V %T :::: %V %T\n", reply, userCode)
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
			"error": "",
		}
	}

	c.ServeJSON()
}

// CreateVerifyCode : 生成验证码
func (c *LoginController) CreateVerifyCode() {
	//生成随机数 模仿验证码
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(8999) + 1000

	conn, err := redis.Dial("tcp", ":6379")
	defer conn.Close()
	if err != nil {
		beego.Info(err)
		return
	}

	conn.Do("set", "codeVerify", code)

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
