var correspondPar = map[byte]byte{
	'(' : ')',
	'[' : ']',
	'{' : '}',
}

func isValid(s string) bool {
    stack := make([]byte,0)
	n := len(s)
	for i := 0; i < n; i++{
		p := s[i]
		if _, existed := correspondPar[p]; existed{
			stack = append(stack, p)
		} else {
			if len(stack) == 0 {
				return false
			}
			topP := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if correspondPar[topP] != p{
				return false
			}
		}
	}
	return len(stack) == 0
}
