func isValid(s string)bool{
	if len(s) == 0 || len(s) >= 3 {return false}
	if s[0] == '0' {return false}
	num, _ := strconv.Atoi(s)
	return num >= 1 && num <= 26
}

func numDecodings(s string) int {
	n := len(s)
	cache := make(map[int]int)
	var numOfDecodeWayFromIdx func(idx int) int
	numOfDecodeWayFromIdx = func(idx int) int{
		if idx >= n {
			return 1
		}
		if val, existed := cache[idx]; existed{
			return val
		}
		result := 0
		if idx < n && isValid(s[idx:idx+1]){
			result += numOfDecodeWayFromIdx(idx+1)
		}
		if idx < n-1 && isValid(s[idx:idx+2]){
			result += numOfDecodeWayFromIdx(idx+2)
		}
		cache[idx]=result
		return result
	}
	return numOfDecodeWayFromIdx(0)
}
