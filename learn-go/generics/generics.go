package main

func getLast[T any](s []T) T {
	if len(s) <= 0 {
		var zeroVar T
		return zeroVar
	}
	return s[len(s)-1]
}
