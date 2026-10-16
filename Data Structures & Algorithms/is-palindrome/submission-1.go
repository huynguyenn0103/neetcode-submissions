func isValid (c byte) bool{
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c<='9')
}

func getLowercase(c byte)byte{
	if c >= 'a' && c <= 'z'{
		return c
	} else if c >= 'A' && c <= 'Z'{
		return c + ('a'-'A')
	} else {
		return c
	}
}

func isPalindrome(s string) bool {
	n := len(s)
	left, right := 0, n-1
	for left <= right{
		for !isValid(s[left]){
			left++
			if left >= n {
				break
			}
		}
		if left >= n{break}
		for !isValid(s[right]){
			right--
			if right < 0{
				break
			}
		}
		if right < 0{break}
		if getLowercase(s[left]) != getLowercase(s[right]){return false}
		left++
		right--
	}
	return true
}
