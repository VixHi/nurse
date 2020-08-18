package models

// Hospital : 医院
type Hospital struct {
	Id            int      `json:"id"`
	HospitalName  string   `json:"hospitalName"`
	HospitalImage string   `json:"hospitalImage"`
	Lat           string   `json:"lat"`
	Lon           string   `json:"lon"`
	Level         string   `json:"level"`
	Nurses        []*Nurse `orm:"reverse(many)";json:"nurses"` //反向
	News          []*News  `orm:"reverse(many)"`               //多对对
}
