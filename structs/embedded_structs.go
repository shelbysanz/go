package main

// Go is not object oriented, but embedded struct provides a data-only inheritance
type sender struct {
	user
	rateLimit int
}

type user struct {
	name   string
	number int
}
