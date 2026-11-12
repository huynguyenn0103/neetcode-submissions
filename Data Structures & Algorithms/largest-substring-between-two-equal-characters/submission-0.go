func maxLengthBetweenEqualCharacters(s string) int {
    result := -1
	mapCharToFirstIdx := make(map[byte]int)
	n := len(s)
	for i := 0; i < n; i++ {
		if firstIdx, existed := mapCharToFirstIdx[s[i]]; existed{
			result = max(result, i - firstIdx - 1)
		} else {
			mapCharToFirstIdx[s[i]] = i
		}
	}
	return result
}