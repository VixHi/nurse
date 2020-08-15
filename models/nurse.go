package models

//Nurse : 护士类
type Nurse struct {
	Id       int
	Name     string
	Phone    string
	Sex      string
	Title    string
	Hospital *Hospital `orm:"rel(fk)"`
}
