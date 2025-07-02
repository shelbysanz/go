package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	messages := [3]string{primary, secondary, tertiary}
	costFirst := len(primary)
	costSecond := costFirst + len(secondary)
	costThird := costSecond + len(tertiary)
	cost := [3]int{costFirst, costSecond, costThird}
	return messages, cost
}
