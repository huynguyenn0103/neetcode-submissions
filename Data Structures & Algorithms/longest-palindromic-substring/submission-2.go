func longestPalindrome(s string) string {
    var final string
	cache := make(map[[2]int]bool)
	n := len(s)
	var isValidPalindrome func(left, right int)bool
	isValidPalindrome = func(left, right int)bool{
		if left > right {
			return true
		}
		key := [2]int{left,right}
		if val, existed := cache[key]; existed{
			return val
		}
		result := false
		if s[left] == s[right]{
			result = isValidPalindrome(left+1,right-1) 
		}
		if result && right - left + 1 > len(final){
			final = s[left:right+1]
		}
		cache[key]=result
		return result
	}
	for i := 0; i < n; i++{
		for j := i; j < n; j++ {
			isValidPalindrome(i,j)
		}
	}
	return final
}
