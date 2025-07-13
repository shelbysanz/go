package main

func maxMessages(thresh int) int {
	const initialCost int = 100
	const additionalFee int = 1
	var totalCost int
	for i := 0; ; i++ {
		totalCost = totalCost + initialCost + (additionalFee * i)
		if totalCost > thresh {
			return i
		}
	}
}
