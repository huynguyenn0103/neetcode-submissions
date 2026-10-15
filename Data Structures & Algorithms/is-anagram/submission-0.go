func isAnagram(s string, t string) bool {
	if len(s) != len(t){return false}
	mapCharToFreq := make(map[byte]int)
	for i := 0; i < len(s); i++{
		mapCharToFreq[s[i]]++
	}
	for i := 0; i < len(t); i++{
		mapCharToFreq[t[i]]--
	}
	for _, freq := range mapCharToFreq{
		if freq != 0 {
			return false
		}
	}
	return true
}
