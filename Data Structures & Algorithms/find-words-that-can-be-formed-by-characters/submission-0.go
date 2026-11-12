import "maps"

func countCharacters(words []string, chars string) int {
    mapCharToFreq := make(map[byte]int)
	n := len(chars)
	for i := 0; i < n; i++{
		mapCharToFreq[chars[i]]++
	}
	result := 0
	for _, word := range words{
		clone := maps.Clone(mapCharToFreq)
		m := len(word)
		isGood := true
		for i := 0; i < m; i++{
			clone[word[i]]--
			if clone[word[i]] < 0 {
				isGood = false
			}
		}
		if isGood{
			result += m
		}
	}
	return result
}