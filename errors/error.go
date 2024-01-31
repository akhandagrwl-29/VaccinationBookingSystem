package errors

import "errors"

var (
	NoInputErr                        = errors.New("No input provided")
	ErrFileOpen                       = errors.New("error in opening file")
	InvalidCommand                    = errors.New("invalid command")
	ErrUserAlreadyExists              = errors.New("user already exists")
	ErrUserAlreadyVaccinated          = errors.New("user already vaccinated")
	ErrSlotNotAvailable               = errors.New("slot not available")
	ErrVaccinationCentreAlreadyExists = errors.New("vaccination centre already exists")
	ErrInvalidBooking                 = errors.New("invalid booking")
)
