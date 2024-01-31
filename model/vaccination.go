package model

type VaccinationCentre struct {
	CentreId string
	State    string
	District string
	Capacity []Capacity
}

type VaccinationApp struct {
	VaccinationApp []VaccinationCentre
}
