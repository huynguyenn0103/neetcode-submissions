func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	cache := make(map[[2]int]int)
	var lcsFromIdx func(idx1, idx2 int)int
	lcsFromIdx = func(idx1, idx2 int)int{
		if idx1 >= m || idx2 >= n {
			return 0
		}
		key := [2]int{idx1,idx2}
		if val, existed := cache[key]; existed{
			return val
		}
		result := 0
		if text1[idx1] == text2[idx2]{
			result = max(result, 1+lcsFromIdx(idx1+1,idx2+1))
		}
		result = max(result, lcsFromIdx(idx1+1,idx2))
		result = max(result, lcsFromIdx(idx1,idx2+1))
		cache[key]=result
		return result
	}
	return lcsFromIdx(0,0)
}
