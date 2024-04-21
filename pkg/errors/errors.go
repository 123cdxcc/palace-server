package errors

import "github.com/pkg/errors"

func New(message string) error {
	return errors.New(message)
}
