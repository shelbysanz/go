package main

import "fmt"

func sum(nums ...int) int {
	var cost int
	for i := range len(nums) {
		cost += nums[i]
	}
	return cost
}

func main() {
	ints := []int{10, 3, 4, 5, 10} // these are test values
	fmt.Println(sum(ints...))
}
