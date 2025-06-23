package main

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	var messageCharLimit int
	if membershipType == "premium" {
		messageCharLimit = 1000
	} else {
		messageCharLimit = 100
	}

	var newUser User

	newUser.Name = name
	newUser.Type = membershipType
	newUser.MessageCharLimit = messageCharLimit

	return newUser
}

// this function is more concise
func altNewUser(name string, membershipType string) User {
	membership := Membership{Type: membershipType}
	if membershipType == "premium" {
		membership.MessageCharLimit = 1000
	} else {
		membership.MessageCharLimit = 100
	}
	return User{Name: name, Membership: membership}
}
