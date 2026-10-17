func isSuccess(mapCharToFreq map[byte]int)bool{
	for _, val := range mapCharToFreq{
		if val > 0 {
			return false
		}
	}
	return true
}

func minWindow(s string, t string) string {
    mapCharToFreq := make(map[byte]int)
	for i := 0; i < len(t); i++{
		mapCharToFreq[t[i]]++
	}
	result, hasResult := "", false
	n := len(s)
	ptr1, ptr2 := 0, 0
	for ptr2 < n{
		freq, existed := mapCharToFreq[s[ptr2]]
		if existed && freq > 0{
			mapCharToFreq[s[ptr2]]--
		} else {
			if existed {
				mapCharToFreq[s[ptr2]]--
			}
			for ptr1 < n{
				freq1, existed1 := mapCharToFreq[s[ptr1]]
				if existed1 && freq1 >= 0{break}
				if existed1{
					mapCharToFreq[s[ptr1]]++
				}
				ptr1++
			}
		}
		if isSuccess(mapCharToFreq){
			subStr := s[ptr1:ptr2+1]
			if !hasResult {
				result = subStr
				hasResult = true
			} else {
				if len(result) > len(subStr){
					result = subStr
				}
			}
		}
		ptr2++
	}
	return result
}


// CBBA

// ADD BBC
// ADD C B ACCDCBB

// CD B BAZZ CBAB

// s[ptr2] Not Existed || freq[s[ptr2]] == 0{
// 	move ptr1 den freq[s[ptr1]] >= 0
// }

