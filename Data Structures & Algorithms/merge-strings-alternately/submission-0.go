func mergeAlternately(word1 string, word2 string) string {
	n1,n2 := len(word1),len(word2)
	ptr1, ptr2 := 0,0
	turn := 0
	result := make([]byte,0)
	for ptr1 < n1 && ptr2 < n2{
		if turn == 0 {
			result = append(result,word1[ptr1])
			ptr1++
		} else {
			result = append(result, word2[ptr2])
			ptr2++
		}
		turn = 1 - turn
	}
	for ptr1 < n1 {
		result = append(result, word1[ptr1])
		ptr1++
	}
	for ptr2 < n2 {
		result = append(result, word2[ptr2])
		ptr2++
	}
	return string(result)
}
