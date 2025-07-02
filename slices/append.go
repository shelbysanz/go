package main

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	dayCosts := []float64{}
	for i := range len(costs) {
		if costs[i].day == day {
			dayCosts = append(dayCosts, costs[i].value)
		}
	}
	return dayCosts
}
