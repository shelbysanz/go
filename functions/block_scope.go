package main

import "fmt"

func splitEmail(email string) (string, string) {
	username, domain := "", ""
	for i, r := range email {
		if r == '@' {
			username = email[:i]
			domain = email[i+1:]
			break
		}
	}
	{
		// this would be another scope
		age := 23
		fmt.Println(age)
	}

	// the following code would not work since it's out of scope
	// fmt.Println(age)
	return username, domain
}
