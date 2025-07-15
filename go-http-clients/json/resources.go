package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
)

func getResources(url string) ([]map[string]any, error) {
	var resources []map[string]any

	res, err := http.Get(url)
	if err != nil {
		return resources, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var unmarshalled []map[string]any
	if err := json.Unmarshal(data, &unmarshalled); err != nil {
		return nil, err
	}

	return unmarshalled, nil

}

func logResources(resources []map[string]any) {
	var formattedStrings []string

	for i := range resources {
		for key, value := range resources[i] {
			formattedStrings = append(formattedStrings, fmt.Sprintf("Key: %s - Value: %v", key, value))
		}
	}

	sort.Strings(formattedStrings)

	for _, str := range formattedStrings {
		fmt.Println(str)
	}
}
