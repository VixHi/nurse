package models

type Hospital struct {
	Id            int
	HospitalName  string
	HospitalImage string
	Lat           string
	Lon           string
	Level         string
	Nurses        []*Nurse `orm:"reverse(many)"`
}
