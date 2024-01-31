package processor

import (
	"VaccinationBookingSystem/controller"
	"VaccinationBookingSystem/errors"
	"VaccinationBookingSystem/model"
	"VaccinationBookingSystem/utils"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ProcessFile(fileName string) {
	file, err := os.Open(fileName)
	defer file.Close()

	if err != nil {
		panic(errors.ErrFileOpen)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputText := scanner.Text()

		if controller.Store == nil {
			controller.InitStore()
			controller.InitUserDetails()
			controller.InitVaccinationApp()
		}

		// add command based on input text
		command, err := addCommand(inputText)
		if err != nil {
			panic(err)
		}

		err = command.Connection.Execute(&command)
		if err != nil {
			panic(err)
		}

	}
}

func addCommand(inputText string) (model.Command, error) {
	var command model.Command

	inputCommand := strings.Split(inputText, " ")
	if len(inputCommand) == 0 {
		return command, errors.InvalidCommand
	}

	command.CommandName = inputCommand[0]

	if command.CommandName == utils.CommandExit {
		os.Exit(1)
	}

	err := addConnection(&command)
	if err != nil {
		return command, err
	}

	err = addArguments(&command, inputText)
	if err != nil {
		return command, err
	}

	return command, nil
}

func addConnection(command *model.Command) error {
	switch command.CommandName {
	case utils.CommandAddUser:
		command.Connection = controller.Store.AddUser()
	case utils.CommandAddVaccinationCentre:
		command.Connection = controller.Store.AddVaccinationCentre()
	case utils.CommandAddCapacity:
		command.Connection = controller.Store.AddCapacity()
	case utils.CommandBookVaccination:
		command.Connection = controller.Store.BookVaccination()
	case utils.CommandListCentres:
		command.Connection = controller.Store.ListVaccinationCentre()
	case utils.CommandListAllBookings:
		command.Connection = controller.Store.ListAllBookings()
	case utils.CommandCancelBooking:
		command.Connection = controller.Store.CancelBooking()
	default:
		return errors.InvalidCommand
	}
	return nil
}

func addArguments(command *model.Command, inputText string) error {
	inputCommand := strings.Split(inputText, " ")

	switch command.CommandName {
	case utils.CommandAddUser:
		if len(inputCommand) != 7 {
			return errors.InvalidCommand
		}
		age, err := strconv.Atoi(inputCommand[4])
		if err != nil {
			return errors.InvalidCommand
		}

		command.Arguments = model.User{
			Id:       inputCommand[1],
			Name:     inputCommand[2],
			Gender:   inputCommand[3],
			Age:      int64(age),
			State:    inputCommand[5],
			District: inputCommand[6],
		}
	case utils.CommandAddVaccinationCentre:
		if len(inputCommand) != 4 {
			return errors.InvalidCommand
		}
		command.Arguments = model.VaccinationCentre{
			CentreId: inputCommand[3],
			State:    inputCommand[1],
			District: inputCommand[2],
			Capacity: []model.Capacity{},
		}
	case utils.CommandAddCapacity:
		if len(inputCommand) != 4 {
			return errors.InvalidCommand
		}

		capacityInUnits, err := strconv.Atoi(inputCommand[3])
		if err != nil {
			return errors.InvalidCommand
		}

		day, err := strconv.Atoi(inputCommand[2])
		if err != nil {
			return errors.InvalidCommand
		}

		command.Arguments = model.Capacity{
			CapacityInUnits: int64(capacityInUnits),
			Day:             int64(day),
			CentreId:        inputCommand[1],
		}
	case utils.CommandBookVaccination:
		if len(inputCommand) != 4 {
			return errors.InvalidCommand
		}
		day, err := strconv.Atoi(inputCommand[2])
		if err != nil {
			return errors.InvalidCommand
		}

		command.Arguments = model.BookVaccination{
			CentreId: inputCommand[1],
			Day:      int64(day),
			UserId:   inputCommand[3],
		}

	case utils.CommandListCentres:
		if len(inputCommand) != 2 {
			return errors.InvalidCommand
		}

		command.Arguments = model.ListVaccinationCentre{
			District: inputCommand[1],
		}

	case utils.CommandListAllBookings:
		if len(inputCommand) != 3 {
			return errors.InvalidCommand
		}

		day, err := strconv.Atoi(inputCommand[2])
		if err != nil {
			return errors.InvalidCommand
		}

		command.Arguments = model.ListAllBookings{
			CentreId: inputCommand[1],
			Day:      int64(day),
		}
	case utils.CommandCancelBooking:
		if len(inputCommand) != 4 {
			return errors.InvalidCommand
		}
		day, err := strconv.Atoi(inputCommand[2])
		if err != nil {
			return errors.InvalidCommand
		}

		command.Arguments = model.CancelBooking{
			CentreId: inputCommand[1],
			Day:      int64(day),
			UserId:   inputCommand[3],
		}

	default:
		return errors.InvalidCommand
	}
	return nil
}
