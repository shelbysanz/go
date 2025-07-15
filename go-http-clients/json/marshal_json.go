package main

import (
	"encoding/json"
	"errors"
)

func marshalAll[T any](items []T) ([][]byte, error) {
	var marshalled [][]byte
	for _, item := range items {
		data, err := json.Marshal(item)
		if err != nil {
			return nil, errors.New("Unable to marshal the data")
		}
		marshalled = append(marshalled, data)
	}

	return marshalled, nil
}
