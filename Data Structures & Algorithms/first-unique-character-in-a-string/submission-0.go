func firstUniqChar(s string) int {
	mapCharToFreq := make(map[byte]int)
	n := len(s)
	for i := 0; i < n; i++ {
		mapCharToFreq[s[i]]++
	}
	for i := 0; i < n; i++ {
		if mapCharToFreq[s[i]] == 1{
			return i
		}
	}
	return -1
}
