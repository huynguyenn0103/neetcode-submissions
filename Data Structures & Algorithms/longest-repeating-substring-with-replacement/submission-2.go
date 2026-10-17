func characterReplacement(s string, k int) int {
	mapOfChars := make(map[byte]bool)
	n := len(s)
	for i := 0; i < n; i++ {
		mapOfChars[s[i]]=true
	}

	result := 0

	for c := range mapOfChars{
		ptr1, ptr2 := 0, 0
		posQueue := make([]int, 0)
		for ptr2 < n {
			if s[ptr2] != c {
				if len(posQueue) == k {
					result = max(result, ptr2-ptr1)
					if k == 0 {
						ptr1 = ptr2+1
					} else {
						front := posQueue[0]
						posQueue = posQueue[1:]
						posQueue = append(posQueue, ptr2)
						ptr1 = front + 1
					}
				} else {
					posQueue = append(posQueue, ptr2)
				}
			}
			ptr2++
		}
		result = max(result, ptr2-ptr1)
	}
	return result
}
