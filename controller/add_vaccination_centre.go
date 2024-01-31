package controller

import (
	"VaccinationBookingSystem/errors"
	"VaccinationBookingSystem/model"
	"fmt"
)

type addVaccinationCentre struct {
	*store
}

func NewAddVaccinationCentre(st *store) *addVaccinationCentre {
	return &addVaccinationCentre{st}
}

func (s *addVaccinationCentre) Execute(command *model.Command) error {

	args, ok := command.Arguments.(model.VaccinationCentre)

	if !ok {
		return errors.InvalidCommand
	}

	IsCentreExists := isCentreAlreadyExists(args.CentreId)
	if IsCentreExists {
		return errors.ErrVaccinationCentreAlreadyExists
	}

	VaccinationApp.VaccinationApp = append(VaccinationApp.VaccinationApp, model.VaccinationCentre{
		CentreId: args.CentreId,
		State:    args.State,
		District: args.District,
	})

	fmt.Println("The current centres are :", VaccinationApp)

	return nil
}

func isCentreAlreadyExists(id string) bool {
	for _, centre := range VaccinationApp.VaccinationApp {
		if centre.CentreId == id {
			return true
		}
	}
	return false
}
