package model

type User struct {
	Id                  string
	Name                string
	Gender              string
	Age                 int64
	State               string
	District            string
	IsVaccinated        bool
	VaccinationDay      int64
	VaccinationCentreId string
}

type UserDetails struct {
	UserDetails []User
}
