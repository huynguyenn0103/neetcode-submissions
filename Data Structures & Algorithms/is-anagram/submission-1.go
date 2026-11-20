func isAnagram(s string, t string) bool {
 m,n := len(s),len(t)
 if m != n {return false}
 mapCharToFreq := make(map[byte]int)
 for i := 0; i < m; i++{
	c := s[i]
	mapCharToFreq[c]++
 }
 for i := 0; i < m; i++{
	c := t[i]
	if mapCharToFreq[c] == 0{
		return false
	}
	mapCharToFreq[c]--
	if mapCharToFreq[c] == 0{
		delete(mapCharToFreq,c)
	}
 }
 return len(mapCharToFreq) == 0
}
