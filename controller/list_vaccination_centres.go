package controller

import (
	"VaccinationBookingSystem/errors"
	"VaccinationBookingSystem/model"
	"fmt"
)

type listVaccinationCentres struct {
	*store
}

func NewListVaccinationCentres(st *store) *listVaccinationCentres {
	return &listVaccinationCentres{st}
}

func (s *listVaccinationCentres) Execute(command *model.Command) error {

	args, ok := command.Arguments.(model.ListVaccinationCentre)

	if !ok {
		return errors.InvalidCommand
	}
	fmt.Printf("The vaccination centres for district: %s\n", args.District)
	for _, centre := range VaccinationApp.VaccinationApp {
		if centre.District == args.District {
			fmt.Println(centre)
		}
	}

	return nil
}
