var mapNumToChars = map[byte][]byte{
	'2': []byte{'a','b','c'},
	'3': []byte{'d','e','f'},
	'4': []byte{'g','h','i'},
	'5': []byte{'j','k','l'},
	'6': []byte{'m','n','o'},
	'7': []byte{'p','q','r','s'},
	'8': []byte{'t','u','v'},
	'9': []byte{'w','x','y','z'},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {return []string{}}
	result := make([]string, 0)
	var backtrack func(curIdx int, curString []byte)
	backtrack = func(curIdx int, curString []byte){
		if curIdx == len(digits){
			result = append(result, string(curString))
			return
		}
		c := digits[curIdx]
		chars := mapNumToChars[c]
		for i := 0; i < len(chars); i++ {
			curString = append(curString, chars[i])
			backtrack(curIdx+1, curString)
			curString = curString[:len(curString)-1]
		}
	}
	backtrack(0, []byte{})
	return result
}
