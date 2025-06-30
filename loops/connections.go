package main

import "fmt"

func countConnections(groupSize int) int {
	var totalConnections int
	for i := groupSize; i > 0; i-- {
		totalConnections = totalConnections + i - 1
	}
	return totalConnections
}

func main() {
	fmt.Println(countConnections(5))
}
