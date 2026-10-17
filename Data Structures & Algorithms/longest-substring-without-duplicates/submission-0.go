func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n <= 1{
		return n
	}
	mapCharToPosition := make(map[byte]int)
	ptr1, ptr2 := 0, 1
	mapCharToPosition[s[0]]=0
	result := 0
	for ptr2 < n {
		c := s[ptr2]
		if pos, existed := mapCharToPosition[c]; existed {
			result = max(result, ptr2-ptr1)
			for ptr1 <= pos{
				delete(mapCharToPosition, s[ptr1])
				ptr1++
			}
		}
		mapCharToPosition[c] = ptr2
		ptr2++
	}
	return max(result, ptr2-ptr1)
}
