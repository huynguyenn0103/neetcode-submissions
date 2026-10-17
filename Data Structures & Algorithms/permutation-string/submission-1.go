func checkInclusion(s1 string, s2 string) bool {
	mapCharToFreq := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		c := s1[i]
		mapCharToFreq[c]++
	}

	n := len(s2)
	ptr1, ptr2 := 0, 0
	for ptr2 < n{
		c := s2[ptr2]
		if freq, existed := mapCharToFreq[c]; existed {
			if freq > 0 {
				mapCharToFreq[c]--
				isFinished := true
				for _, freq := range mapCharToFreq{
					if freq != 0{
						isFinished = false
					}
				}
				if isFinished{return true}
			} else {
				for s2[ptr1] != s2[ptr2]{
					mapCharToFreq[s2[ptr1]]++
					ptr1++
				}
				ptr1++
			}
		} else {
			for ptr1 < ptr2 {
				mapCharToFreq[s2[ptr1]]++
				ptr1++
			}
			ptr1 = ptr2 + 1
		}
		ptr2++
	}
	return false
}
