func longestPalindrome(s string) int {
    result := 0
	mapCharToFreq := make(map[byte]int)
	n := len(s)
	for i := 0; i < n; i++ {
		mapCharToFreq[s[i]]++
	}
	for _, val := range mapCharToFreq{
		if val % 2 == 0 {
			result = result + val
		} else {
			result = result + val -1
		}
	}
	if result < n {
		result = result + 1
	}
	return result
}