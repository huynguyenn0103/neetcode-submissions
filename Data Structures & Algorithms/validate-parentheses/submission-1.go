var mapOpenToCloseP = map[byte]byte{
	')' : '(',
	'}' : '{',
	']' : '[',
}

func isValid(s string) bool {
    stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if openP, existed := mapOpenToCloseP[c]; existed{
			if len(stack)==0 || openP != stack[len(stack)-1] {return false}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}
	return len(stack) == 0
}
