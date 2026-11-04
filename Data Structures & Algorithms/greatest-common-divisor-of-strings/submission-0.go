func canDivide(str string, bytes []byte)bool {
	idx := 0
	m, n := len(str),len(bytes)
	for i := 0; i < m; i++ {
		if str[i] != bytes[idx]{return false}
		idx++
		if idx == n{
			idx = 0
		}
	}
	return idx == 0
}

func gcdOfStrings(str1 string, str2 string) string {
	ptr := 0
	largestPrefix := make([]byte,0)
	for ptr < len(str1) && ptr < len(str2){
		if str1[ptr] == str2[ptr]{
			largestPrefix = append(largestPrefix,str1[ptr])
			ptr++
		} else {
			break
		}
	}
	n := len(largestPrefix)
	for end := n; end > 0; end--{
		bytes := largestPrefix[0:end]
		if canDivide(str1,bytes) && canDivide(str2,bytes){
			return string(bytes)
		}
	}
	return ""
}
