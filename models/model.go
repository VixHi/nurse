package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(new(Image), new(Nurse))
	orm.RunSyncdb("default", false, true)
}
