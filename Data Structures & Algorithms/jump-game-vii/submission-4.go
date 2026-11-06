// 0,0
// 2 -> 5
// 4 -> 7 && 7 -> 10
// 6 -> 9 && 9 -> 12 && 12 -> 15

func canReach(s string, minJump int, maxJump int) bool {
	minPos, maxPos := 0,0
	n := len(s)
	for maxPos < n-1 {
		minPos = minPos + minJump
		maxPos = maxPos + maxJump
		for minPos < maxPos && minPos < n && s[minPos] != '0'{
			minPos++
		}
		for maxPos > minPos && maxPos < n && s[maxPos] != '0' {
			maxPos--
		}
		if minPos < n && maxPos < n && (s[minPos] == '1' || s[maxPos] == '1'){
			return false
		}
	}
	return minPos == n-1 || minPos + minJump <= n-1 && s[n-1] == '0'
}

//011000001011110000
//14,21
