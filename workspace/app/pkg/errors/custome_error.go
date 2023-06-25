package errors

import (
	"errors"
	// "time"
)

func ErrorFromDao(err error) error {
	if err != nil {
		return message(err, "error from dao")
	}
	return nil
}

func message(err error, message string) error {
	//TODO: change message
	return errors.New(message)
}
