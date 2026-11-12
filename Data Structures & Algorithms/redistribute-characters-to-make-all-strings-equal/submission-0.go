func makeEqual(words []string) bool {
    n := len(words)
	mapCharToFreq := make(map[byte]int)
	for _, word := range words{
		m := len(word)
		for i := 0; i < m; i++ {
			mapCharToFreq[word[i]]++
		}
	}
	for _, val := range mapCharToFreq{
		if val % n != 0 {
			return false
		}
	}
	return true
}