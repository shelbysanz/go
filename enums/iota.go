package main

/*
Define an emailStatus type that uses iota syntax to represent the following states:
	- emailBounced: 0
	- emailInvalid: 1
	- emailDelivered: 2
	- emailOpened: 3
*/

type emailStatus int

const (
	emailBounced emailStatus = iota
	emailInvalid
	emailDelivered
	emailOpened
)
