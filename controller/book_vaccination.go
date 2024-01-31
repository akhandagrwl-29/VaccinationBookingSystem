package controller

import (
	"VaccinationBookingSystem/errors"
	"VaccinationBookingSystem/model"
)

type bookVaccination struct {
	*store
}

func NewBookVaccination(st *store) *bookVaccination {
	return &bookVaccination{st}
}

func (s *bookVaccination) Execute(command *model.Command) error {

	args, ok := command.Arguments.(model.BookVaccination)

	if !ok {
		return errors.InvalidCommand
	}

	isUserVaccinated := isUserAlreadyVaccinated(args.UserId)
	if isUserVaccinated {
		return errors.ErrUserAlreadyVaccinated
	}

	isSlotAvailable := isSlotAvailable(args)
	if !isSlotAvailable {
		return errors.ErrSlotNotAvailable
	}

	bookVaccinationAndUpdateSlot(args)

	// fmt.Println("The current users are :", UserDetails)

	return nil
}

func isUserAlreadyVaccinated(userId string) bool {
	for _, user := range UserDetails.UserDetails {
		if user.Id == userId {
			if user.IsVaccinated {
				return true
			}
		}
	}
	return false
}

func isSlotAvailable(args model.BookVaccination) bool {

	for _, centre := range VaccinationApp.VaccinationApp {
		if centre.CentreId == args.CentreId {
			for _, capacity := range centre.Capacity {
				if capacity.Day == args.Day && capacity.CapacityInUnits > 0 {
					return true
				}
			}
		}
	}
	return false
}

func bookVaccinationAndUpdateSlot(args model.BookVaccination) {
	for userIdx, user := range UserDetails.UserDetails {
		if user.Id == args.UserId {
			UserDetails.UserDetails[userIdx].IsVaccinated = true
			UserDetails.UserDetails[userIdx].VaccinationDay = args.Day
			UserDetails.UserDetails[userIdx].VaccinationCentreId = args.CentreId
		}
	}

	for centreIdx, centre := range VaccinationApp.VaccinationApp {
		if centre.CentreId == args.CentreId {
			for capacityIdx, capacity := range centre.Capacity {
				if capacity.Day == args.Day && capacity.CapacityInUnits > 0 {
					VaccinationApp.VaccinationApp[centreIdx].Capacity[capacityIdx].CapacityInUnits -= 1
				}
			}
		}
	}

}
