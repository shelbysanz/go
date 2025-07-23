package main

import "fmt"

/*
	Explanation

	Rules:
	a ^ 0 = a
	a ^ a = 0

	Logic:
	1 ^ 2 ^ n-1 = combination of bit representation
	1 ^ 2 ^ missing number = combination of bit represenation (without a number)

	Combining rules and logic:
	Since its comparing them, there is a missing value,
	by the rule of a ^ 0 = a, then the missing value is returned
*/

func xorArr(arr []int) int {

	var elem int
	for i := range arr {
		elem ^= arr[i]
	}
	return elem
}

func main() {
	res := xorArr([]int{1, 2, 3, 4, 5})
	res2 := xorArr([]int{1, 2, 4, 5})
	fmt.Println(res ^ res2)
}
