import "slices"
var mapDigitToChar = map[byte][]byte{
	'2' : []byte{'a','b','c'},
	'3' : []byte{'d','e','f'},
	'4' : []byte{'g','h','i'},
	'5' : []byte{'j','k','l'},
	'6' : []byte{'m','n','o'},
	'7' : []byte{'p','q','r','s'},
	'8' : []byte{'t','u','v'},
	'9' : []byte{'w','x','y','z'},
}

func letterCombinations(digits string) []string {
	combination := make([]byte,0)
	result := make([]string,0)
	n := len(digits)
	var buildCombinationFromIdx func(idx int)
	buildCombinationFromIdx = func(idx int){
		if idx >= n {
			if len(combination) > 0 {
				result = append(result, string(slices.Clone(combination)))
			}
			return
		}
		c := digits[idx]
		for _, char := range mapDigitToChar[c]{
			combination = append(combination, char)
			buildCombinationFromIdx(idx+1)
			combination = combination[:len(combination)-1]
		}
	}
	buildCombinationFromIdx(0)
	return result
}
