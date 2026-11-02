func validPalindrome(s string) bool {
	n := len(s)
	var check func(left, right, numOfDelete int)bool
	check = func(left, right, numOfDelete int)bool{
		if left > right {return true}
		if s[left] == s[right] {
			return check(left+1,right-1,numOfDelete)
		} else {
			if numOfDelete == 1{return false}
			nextLeftOk := left <= n-2 && s[left+1] == s[right]
			nextRightOk := right >= 1 && s[left] == s[right-1]
			if nextLeftOk && nextRightOk {
				return check(left+1,right,numOfDelete+1) || check(left,right-1,numOfDelete+1)
			} else if nextLeftOk {
				return check(left+1,right,numOfDelete+1)
			} else if nextRightOk{
				return check(left,right-1,numOfDelete+1)
			} else {
				return false
			}
		}
	}
	return check(0,len(s)-1,0)
}
