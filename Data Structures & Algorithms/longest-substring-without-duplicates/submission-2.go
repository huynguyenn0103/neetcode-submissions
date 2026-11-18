func lengthOfLongestSubstring(s string) int {
	mapCharToExisted := make(map[byte]bool)
	left := 0
	n := len(s)
	result := 0
	for right := 0; right < n; right++ {
		c := s[right]
		if mapCharToExisted[c] == false{
			mapCharToExisted[c] = true
			continue
		}
		result = max(result, right - left)
		for s[left] != c{
			mapCharToExisted[s[left]]=false
			left++
		}
		left++
	}
	return max(result, n - left)
}
