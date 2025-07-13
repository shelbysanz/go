package main

func getNameCounts(names []string) map[rune]map[string]int {
	nameCounts := map[rune]map[string]int{}

	for _, name := range names {
		nameRune := []rune(name)[0]
		/*
		 Making this more concise and removing unecessary variable
		*/
		if _, ok := nameCounts[nameRune]; !ok {
			nameCounts[nameRune] = map[string]int{}
		}
		/*
		 Initilizing the name map with 0 is unnecesarry since the default,
		 value returned is 0 if the key does not exist, so I can just increment
		*/
		nameCounts[nameRune][name]++
	}
	return nameCounts
}
