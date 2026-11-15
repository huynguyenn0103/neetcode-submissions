import "maps"

func checkInclusion(s1 string, s2 string) bool {
	m, n := len(s1), len(s2)
	if m > n {return false}
	mapCharInS1 := make(map[byte]int)
	for i := 0; i < m; i++ {
		mapCharInS1[s1[i]]++
	}
	for i := 0; i < n; i++{
		c := s2[i]
		if mapCharInS1[c] > 0 {
			cloneMapS1 := maps.Clone(mapCharInS1)
			for j := i; j < n; j++ {
				curChar := s2[j]
				if cloneMapS1[curChar] == 0 {
					break
				}
				cloneMapS1[curChar]--
				if cloneMapS1[curChar] == 0{
					delete(cloneMapS1, curChar)
					if len(cloneMapS1) == 0{
						return true
					}
				}
			}
		}
	}
	return false
}
