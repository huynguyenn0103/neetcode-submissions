func wordBreak(s string, wordDict []string) bool {
    cacheResultFromIdx := make(map[int]bool)
	mapDict := make(map[string]bool)
	for _, word := range wordDict{
		mapDict[word]=true
	}
	n := len(s)
	var canBreakFromIdx func(idx int)bool
	canBreakFromIdx = func(idx int)bool{
		if idx >= n {
			return true
		}
		if result, existed := cacheResultFromIdx[idx]; existed{
			return result
		}
		for i := idx+1; i <= n; i++{
			str := s[idx:i]
			if mapDict[str]{
				if canBreakFromIdx(i){
					return true
				}
			}
		}
		cacheResultFromIdx[idx]=false
		return false
	}
	return canBreakFromIdx(0)
}
