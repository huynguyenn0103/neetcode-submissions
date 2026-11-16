func isValid(c byte)bool{
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func toLowercase(c byte)byte{
	if c >= 'a' && c <= 'z' {return c}
	return 'a' + (c - 'A')
}


func isPalindrome(s string) bool {
	n := len(s)
	left, right := 0, n-1
	for left <= right{
		if !isValid(s[left]){
			left++
			continue
		}
		if !isValid(s[right]){
			right--
			continue
		}
		charLeft := toLowercase(s[left])
		charRight := toLowercase(s[right])
		if charLeft != charRight{
			return false
		}
		left++
		right--
	}
	return true
}
