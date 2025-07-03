package main

import (
	"errors"
)

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}
	userMap := make(map[string]user)
	for i, name := range names {
		var user user
		user.name = name
		user.phoneNumber = phoneNumbers[i]
		userMap[name] = user
	}
	return userMap, nil
}

type user struct {
	name        string
	phoneNumber int
}
