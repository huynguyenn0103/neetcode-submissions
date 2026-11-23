func countSubstrings(s string) int {
	n := len(s)
    var extend func(left, right int)int
	extend = func(left, right int)int{
		result := 0
		for left >= 0 && right < n && s[left]==s[right]{
			result++
			left--
			right++
		}
		return result
	}
	result := 0
	for i := 0; i < n; i++ {
		result = result + 1 + extend(i-1,i+1)
		result = result + extend(i-1,i)
	}
	return result
}
