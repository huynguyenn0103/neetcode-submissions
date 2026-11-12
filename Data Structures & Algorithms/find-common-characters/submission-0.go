func createMapFromWord(word string) map[byte]int{
	n := len(word)
	mapChar := make(map[byte]int)
	for i := 0; i < n; i++{
		mapChar[word[i]]++
	}
	return mapChar
}

func combineMap(map1, map2 map[byte]int){
	for key, _ := range map1{
		map1[key] = min(map1[key], map2[key])
	}
}

func commonChars(words []string) []string {
    resultMap := createMapFromWord(words[0])
	n := len(words)
	for i := 1; i < n; i++ {
		combineMap(resultMap, createMapFromWord(words[i]))
	}
	result := make([]string, 0)
	for key, val := range resultMap{
		for val > 0 {
			result = append(result, string(key))
			val--
		}
	}
	return result
}