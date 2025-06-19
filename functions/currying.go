package main

import (
	"errors"
	"fmt"
)

/*
 - Partial application: Can partially apply some arguments
 - Function factory: Can create specialized functions from general ones
 - Main benefits:
	- Reusability, trying to follow DRY as much as possible
	- Composition, build complex behavior by combining more simple functions
	- Configuration, can preconfigure functions with specific behaviors
*/

// getLogger takes a function that formats two strings into
// a single string and returns a function that formats two strings but prints
// the result instead of returning it
func getLogger(formatter func(string, string) string) func(string, string) {
	return func(x string, y string) {
		results := formatter(x, y)
		fmt.Println(results);
	}
}

// don't touch below this line

func test(first string, errors []error, formatter func(string, string) string) {
	defer fmt.Println("====================================")
	logger := getLogger(formatter)
	fmt.Println("Logs:")
	for _, err := range errors {
		logger(first, err.Error())
	}
}

func colonDelimit(first, second string) string {
	return first + ": " + second
}
func commaDelimit(first, second string) string {
	return first + ", " + second
}

func main() {
	dbErrors := []error{
		errors.New("out of memory"),
		errors.New("cpu is pegged"),
		errors.New("networking issue"),
		errors.New("invalid syntax"),
	}
	test("Error on database server", dbErrors, colonDelimit)

	mailErrors := []error{
		errors.New("email too large"),
		errors.New("non alphanumeric symbols found"),
	}
	test("Error on mail server", mailErrors, commaDelimit)
}
