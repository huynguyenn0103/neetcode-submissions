func isNum(c byte)bool{
	return c >= '0' && c <= '9'
}

func reverseBytesToNum(bytes []byte)int{
	num := 0
	n := len(bytes)
	for i := n-1; i >= 0; i-- {
		num = num*10 + int(bytes[i]-'0')
	}
	return num
}

func reverseBytesAndDup(bytes []byte, freq int) []byte{
	result := make([]byte,0)
	n := len(bytes)
	for freq > 0 {
		for i := n-1; i >= 0; i--{
			result = append(result, bytes[i])
		}
		freq--
	}
	return result
}

func decodeString(s string) string {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == ']' {
			tmpChars := make([]byte,0)
			for stack[len(stack)-1] != '['{
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				tmpChars = append(tmpChars,top)
			}
			stack = stack[:len(stack)-1]
			tmpNums := make([]byte,0)
			for len(stack) > 0{
				top := stack[len(stack)-1]
				if isNum(top){
					tmpNums = append(tmpNums, top)
					stack = stack[:len(stack)-1]
				} else {
					break
				}
			}
			freq := reverseBytesToNum(tmpNums)
			newChars := reverseBytesAndDup(tmpChars,freq)
			stack = append(stack, newChars...)
		} else {
			stack = append(stack,s[i])
		}
	}
	return string(stack)
}
