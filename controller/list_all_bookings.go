package controller

import (
	"VaccinationBookingSystem/errors"
	"VaccinationBookingSystem/model"
	"fmt"
)

type listAllBookings struct {
	*store
}

func NewListAllBookings(st *store) *listAllBookings {
	return &listAllBookings{st}
}

func (s *listAllBookings) Execute(command *model.Command) error {

	args, ok := command.Arguments.(model.ListAllBookings)

	if !ok {
		return errors.InvalidCommand
	}

	fmt.Printf("Listing all the bookings for centreId: %s and day: %d\n", args.CentreId, args.Day)
	for _, user := range UserDetails.UserDetails {
		if user.VaccinationCentreId == args.CentreId && user.VaccinationDay == args.Day {
			fmt.Printf("booking user id: %s\n", user.Id)
		}
	}

	return nil
}
