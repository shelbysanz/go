package main

func adder() func(int) int {
	var sum int
	return func(value int) int {
		sum += value
		return sum
	}
}

/*
	Recap of Closer
	- Initializing the adder function (its the creator of this behavior)
		- Example: myAdder := adder()
	- I would then call myAdder() to access the state of that sum
	  or of course the myAdder(4) to sum 4 to the running total
*/
