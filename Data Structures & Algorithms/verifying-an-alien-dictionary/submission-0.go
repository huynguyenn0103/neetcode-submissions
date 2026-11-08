

func isAlienSorted(words []string, order string) bool {
	mapByteToIdx := make(map[byte]int)
	for i:= 0; i<len(order); i++ {
		mapByteToIdx[order[i]] = i
	}
	var isValid func(str1, str2 string)bool
	isValid = func(str1, str2 string)bool{
		ptr := 0
		for ptr < len(str1) && ptr < len(str2){
			if str1[ptr] == str2[ptr]{
				ptr++
				continue
			}
			return mapByteToIdx[str1[ptr]] < mapByteToIdx[str2[ptr]]
		}
		return len(str1) <= len(str2)
	}
	for i := 1; i < len(words); i++ {
		if !isValid(words[i-1],words[i]){
			return false
		}
	}
	return true
}
