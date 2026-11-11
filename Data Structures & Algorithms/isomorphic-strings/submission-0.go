func isIsomorphic(s string, t string) bool {
	mapChar := make(map[byte]byte)
	reverseMapChar := make(map[byte]byte)
	if len(s) != len(t){
		return false
	}
	for i := 0; i < len(s); i++ {
		c1 := s[i]
		c2 := t[i]
		if mapping, existed := mapChar[c1]; existed {
			if mapping != c2 {
				return false
			}
		} else {
			if _, existed := reverseMapChar[c2]; existed {
				return false
			} else {
				mapChar[c1] = c2
				reverseMapChar[c2] = c1
			}
		}
	}
	return true
}
