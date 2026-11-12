func maxNumberOfBalloons(text string) int {
    mapCharToFreq := make(map[byte]int)
	n := len(text)
	for i := 0; i < n; i++{
		mapCharToFreq[text[i]]++
	}
	result := 0
	target := "balloon"
	for {
		shouldBreak := false
		for i := 0; i < len(target); i++ {
			mapCharToFreq[target[i]]--
			if mapCharToFreq[target[i]] < 0 {
				shouldBreak = true
				break
			}
		}
		if shouldBreak {
			break
		}
		result++
	}
	return result
}