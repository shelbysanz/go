package main

/*
	for this function range would work best
	but for the sake of using loops, i will portray
	this function in this way
*/

func bulkSend(numMessages int) float64 {
	var totalCost float64
	const additionalFee float64 = 0.01
	for i := 0; i < numMessages; i++ {
		totalCost = totalCost + 1.0 + float64(i)*additionalFee
	}
	return totalCost
}
