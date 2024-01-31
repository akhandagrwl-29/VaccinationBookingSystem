package controller

import (
	"VaccinationBookingSystem/errors"
	"VaccinationBookingSystem/model"
	"fmt"
)

type addCapacity struct {
	*store
}

func NewAddCapcity(st *store) *addCapacity {
	return &addCapacity{st}
}

func (s *addCapacity) Execute(command *model.Command) error {

	args, ok := command.Arguments.(model.Capacity)

	if !ok {
		return errors.InvalidCommand
	}

	for centreIdx, centre := range VaccinationApp.VaccinationApp {
		if args.CentreId == centre.CentreId {
			for capacityIdx, capacity := range centre.Capacity {
				if capacity.Day == args.Day {
					VaccinationApp.VaccinationApp[centreIdx].Capacity[capacityIdx].CapacityInUnits = args.CapacityInUnits
					return nil
				}
			}

			// If no capcity is present for the same day, append a new one
			VaccinationApp.VaccinationApp[centreIdx].Capacity = append(VaccinationApp.VaccinationApp[centreIdx].Capacity, model.Capacity{
				CapacityInUnits: args.CapacityInUnits,
				Day:             args.Day,
				CentreId:        args.CentreId,
			})
		}
	}

	fmt.Println("The current centres with capacity are :", VaccinationApp)

	return nil
}
