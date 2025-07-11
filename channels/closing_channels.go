package main

func countReports(numSentCh chan int) int {
	var count int
	func(count *int) {
		for {
			num, ok := <-numSentCh
			*count += num
			if !ok {
				break
			}
		}
	}(&count)
	return count
}

// don't touch below this line

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}
