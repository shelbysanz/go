package main

import (
	"fmt"
)

func printPrimes(max int) {
	for i := 2; i < max+1; i++ {
		if i == 2 {
			fmt.Println(i)
			continue
		}
		if i%2 == 0 {
			continue
		}
		var isPrimeNumber = true
		for k := 3; k*k <= i; k = k + 2 {
			if i%k == 0 {
				isPrimeNumber = false
				break
			}
		}
		if isPrimeNumber {
			fmt.Println(i)
		}
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
