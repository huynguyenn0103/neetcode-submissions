
func isValid(mapCharToFreq map[byte]int)bool{
	for _, freq := range mapCharToFreq{
		if freq > 0 {return false}
	}
	return true
}

func minWindow(s string, t string) string {
    mapCharToFreq := make(map[byte]int)
	m,n := len(s), len(t)
	for i := 0; i < n; i++ {
		mapCharToFreq[t[i]]++
	}
	result := ""
	l := 0
	for r := 0; r < m; r++{
		c := s[r]
		if _, existed := mapCharToFreq[c]; !existed{
			continue
		} else {
			mapCharToFreq[c]--
			if isValid(mapCharToFreq){
				for {
					freq, existed := mapCharToFreq[s[l]]
					if !existed || freq + 1 <= 0 {
						if existed {
							mapCharToFreq[s[l]]++
						}
						l++
					} else {
						break
					}
				}
				if result == "" || (r-l+1) < len(result){
					result = s[l:r+1]
				}
				mapCharToFreq[s[l]]++
				l++
			}
		}
	}
	return result
}
