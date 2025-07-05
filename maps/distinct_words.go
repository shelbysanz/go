package main

import (
	"strings"
)

func countDistinctWords(messages []string) int {
	distinctWordsStruct := make(map[string]struct{})
	for _, msg := range messages {
		for _, word := range strings.Fields(strings.ToLower(msg)) {
			distinctWordsStruct[word] = struct{}{}
		}
	}
	return len(distinctWordsStruct)
}
