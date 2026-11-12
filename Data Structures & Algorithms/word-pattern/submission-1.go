func wordPattern(pattern string, s string) bool {
    mapPToS := make(map[byte]string)
	mapSToP := make(map[string]byte)
	strs := strings.Split(s, " ")
	m := len(pattern)
	if m != len(strs){
		return false
	}
	for i := 0; i < m; i++ {
		correspondS, existedCorrespondS := mapPToS[pattern[i]]
		correspondP, existedCorrespondP := mapSToP[strs[i]]
		if !existedCorrespondS && !existedCorrespondP{
			mapPToS[pattern[i]] = strs[i]
			mapSToP[strs[i]] = pattern[i]
		} else if existedCorrespondS && existedCorrespondP{
			if correspondS != strs[i] || correspondP != pattern[i]{
				return false
			}
		} else {
			return false
		}
	}
	return true
}