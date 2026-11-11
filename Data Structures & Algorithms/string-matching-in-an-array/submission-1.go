func isSubstring(str1, str2 string)bool {
	if len(str1) > len(str2){
		return false
	}
	ptr1, ptr2 := 0,0
	for ptr1 < len(str1) && ptr2 < len(str2){
		if str1[ptr1] != str2[ptr2]{
			ptr1 = 0
		} else {
			ptr1++
		}
		ptr2++
	}
	return ptr1 == len(str1)
}

func stringMatching(words []string) []string {
    n := len(words)
	result := make([]string,0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {continue}
			if isSubstring(words[i],words[j]){
				result = append(result, words[i])
				break
			}
		}
	}
	return result
}