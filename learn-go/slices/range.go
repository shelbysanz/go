package main

import "slices"

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, message := range msg {
		for k := range badWords {
			if message == badWords[k] {
				return i
			}
		}
	}
	return -1
}

// can also do this:
func indexOfFirstBadWordTwo(msg []string, badWords []string) int {
	for i, message := range msg {
		if slices.Contains(badWords, message) {
			return i
		}
	}
	return -1
}
