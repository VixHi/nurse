package models

import "time"

// News : 新闻
type News struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content`
	CreateTime time.Time `orm:"auto_now",json:"creatTime"`
	Hospital   *Hospital `orm:"rel(fk)"`
	Nurses     []*Nurse  `orm:"reverse(many)"`
}
