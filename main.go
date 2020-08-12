package main

import (
	"fmt"
	"time"
	_ "viv/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
)

func main() {

	fmt.Print("hello world")
	fmt.Println(beego.WorkPath)
	// beego.info("123")
	cac, err := cache.NewCache("redis", `{"conn":"127.0.0.1:6379"}`)
	if err != nil {
		beego.Info(err)
		return
	}

	cac.Put("age", 18, 100*time.Second)

	fmt.Println(cac.Get("age"))

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/b"] = "swagger"
	}
	beego.Run()

}
