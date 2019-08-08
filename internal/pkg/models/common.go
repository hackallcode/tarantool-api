package models

import (
	"errors"
)

var (
	NotFoundError      = errors.New("not found")
	AlreadyExistsError = errors.New("already exists")
	IncorrectDataError = errors.New("incorrect data")
)

type InputModel interface {
	UnmarshalJSON(data []byte) error
	Validate() bool
}

type OutputModel interface {
	MarshalJSON() ([]byte, error)
}
