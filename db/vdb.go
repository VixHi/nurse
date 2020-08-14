package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego/orm"
)

func init() {

	fmt.Println("=======db==========")
	// 参数1        数据库的别名，用来在 ORM 中切换数据库使用
	// 参数2        driverName
	// 参数3        对应的链接字符串
	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)
	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("default", "mysql", "root:HiVix@1009.@tcp(127.0.0.1:3306)/nurse?charset=utf8", maxIdle, maxConn)
}
