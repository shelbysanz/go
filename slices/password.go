package main

func isValidPassword(password string) bool {
	hasUpper := false
	hasLower := false
	if len(password) >= 5 && len(password) < 13 {
		for _, letter := range password {
			if letter >= 65 && letter <= 90 {
				hasUpper = true
			}
			if letter >= 48 && letter <= 57 {
				hasLower = true
			}
		}
		if hasUpper && hasLower {
			return true
		}
	}
	return false
}
