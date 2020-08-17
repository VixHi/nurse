package models

//Nurse : 护士类
type Nurse struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Phone    string    `json:"phone"`
	Sex      string    `json:"sex"`
	Title    string    `json:"title"`
	Hospital *Hospital `orm:"rel(fk)",json:"hospital"`
}
