package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(new(Image), new(Nurse), new(Hospital))
	orm.RunSyncdb("default", false, true)
}
