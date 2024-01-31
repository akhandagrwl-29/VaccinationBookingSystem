package controller

import (
	"VaccinationBookingSystem/errors"
	"VaccinationBookingSystem/model"
	"fmt"
)

type addUser struct {
	*store
}

func NewAddUser(st *store) *addUser {
	return &addUser{st}
}

func (s *addUser) Execute(command *model.Command) error {

	args, ok := command.Arguments.(model.User)

	if !ok {
		return errors.InvalidCommand
	}

	isUserExists := isUserAlreadyExists(args.Id)
	if isUserExists {
		return errors.ErrUserAlreadyExists
	}

	UserDetails.UserDetails = append(UserDetails.UserDetails, model.User{
		Id:           args.Id,
		Name:         args.Name,
		Gender:       args.Gender,
		Age:          args.Age,
		State:        args.State,
		District:     args.District,
		IsVaccinated: false,
	})

	fmt.Println("The current users are :", UserDetails)

	return nil
}

func isUserAlreadyExists(id string) bool {
	for _, user := range UserDetails.UserDetails {
		if user.Id == id {
			return true
		}
	}
	return false
}
