package main

import (
	"errors"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	if plan == planFree {
		return messages[:2], nil
	}
	if plan == planPro {
		return messages[:], nil
	}
	return nil, errors.New("unsupported plan")
}
