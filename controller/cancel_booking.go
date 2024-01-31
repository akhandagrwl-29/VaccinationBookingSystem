package controller

import (
	"VaccinationBookingSystem/errors"
	"VaccinationBookingSystem/model"
	"fmt"
)

type cancelBooking struct {
	*store
}

func NewCancelBooking(st *store) *cancelBooking {
	return &cancelBooking{st}
}

func (s *cancelBooking) Execute(command *model.Command) error {

	args, ok := command.Arguments.(model.CancelBooking)

	if !ok {
		return errors.InvalidCommand
	}

	isUserVaccinated := isUserAlreadyVaccinated(args.UserId)
	if !isUserVaccinated {
		return errors.ErrInvalidBooking
	}

	cancelVaccinationAndUpdateSlot(args)

	fmt.Println("The current users are :", UserDetails)

	return nil
}

func cancelVaccinationAndUpdateSlot(args model.CancelBooking) {
	for userIdx, user := range UserDetails.UserDetails {
		if user.Id == args.UserId {
			UserDetails.UserDetails[userIdx].IsVaccinated = false
			UserDetails.UserDetails[userIdx].VaccinationDay = 0
			UserDetails.UserDetails[userIdx].VaccinationCentreId = ""
		}
	}

	for centreIdx, centre := range VaccinationApp.VaccinationApp {
		if centre.CentreId == args.CentreId {
			for capacityIdx, capacity := range centre.Capacity {
				if capacity.Day == args.Day {
					VaccinationApp.VaccinationApp[centreIdx].Capacity[capacityIdx].CapacityInUnits += 1
				}
			}
		}
	}

}
