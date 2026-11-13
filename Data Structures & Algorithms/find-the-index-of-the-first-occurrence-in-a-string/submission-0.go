func strStr(haystack string, needle string) int {
	m, n := len(haystack), len(needle)
	for i := 0; i < m; i++{
		ptr1 := i
		ptr2 := 0
		for ptr1 < m && ptr2 < n{
			if haystack[ptr1] != needle[ptr2]{
				break
			}
			ptr1++
			ptr2++
		}
		if ptr2 == n{
			return i
		}
	}
	return -1
}
