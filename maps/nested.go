package main

func getNameCounts(names []string) map[rune]map[string]int {
	nameCounts := map[rune]map[string]int{}

	for _, name := range names {
		nameRune := []rune(name)[0]           // b
		runeNames, ok := nameCounts[nameRune] // { billy: 1 }
		if !ok {                              // if it doesn't exist, it initializes a rune key with a map
			nameCounts[nameRune] = map[string]int{}
			runeNames = nameCounts[nameRune]
		}
		_, ok = runeNames[name] // checking if name already exists
		if !ok {
			runeNames[name] = 0 // if it doesn't exist, it initializes the name key with a 0 count
		}
		runeNames[name] += 1
	}
	return nameCounts
}
