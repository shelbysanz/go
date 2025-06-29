package main

import (
	"errors"
)

func validateStatus(status string) error {
	statusLength := len(status)
	if statusLength == 0 {
		return errors.New("status cannot be empty")
	}
	if statusLength > 140 {
		return errors.New("status exceeds 140 characters")
	}
	return nil
}
