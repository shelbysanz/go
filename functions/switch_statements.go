package main

func getMonthlyPrice(tier string) int {
	var price float64
	switch tier {
	case "basic":
		price = 100.00
	case "premium":
		price = 150.00
	case "enterprise":
		price = 500.00
	default:
		price = 0
	}

	pennies := int(price) * 100

	return pennies

}
