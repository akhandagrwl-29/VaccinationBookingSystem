package controller

import "VaccinationBookingSystem/model"

type store struct {
	addUser               model.CMD
	addCapacity           model.CMD
	addVaccinationCentre  model.CMD
	bookVaccination       model.CMD
	cancelBooking         model.CMD
	listVaccinationCentre model.CMD
	listAllBookings       model.CMD
}

func (s *store) AddUser() model.CMD {
	return s.addUser
}

func (s *store) AddVaccinationCentre() model.CMD {
	return s.addVaccinationCentre
}

func (s *store) AddCapacity() model.CMD {
	return s.addCapacity
}

func (s *store) ListVaccinationCentre() model.CMD {
	return s.listVaccinationCentre
}

func (s *store) BookVaccination() model.CMD {
	return s.bookVaccination
}

func (s *store) ListAllBookings() model.CMD {
	return s.listAllBookings
}

func (s *store) CancelBooking() model.CMD {
	return s.cancelBooking
}

var Store *store
var VaccinationApp *model.VaccinationApp
var UserDetails *model.UserDetails

func InitVaccinationApp() {
	VaccinationApp = new(model.VaccinationApp)
}

func InitUserDetails() {
	UserDetails = new(model.UserDetails)
}

func InitStore() {
	Store = new(store)
	Store.addUser = NewAddUser(Store)
	Store.addVaccinationCentre = NewAddVaccinationCentre(Store)
	Store.addCapacity = NewAddCapcity(Store)
	Store.bookVaccination = NewBookVaccination(Store)
	Store.listVaccinationCentre = NewListVaccinationCentres(Store)
	Store.listAllBookings = NewListAllBookings(Store)
	Store.cancelBooking = NewCancelBooking(Store)
}
