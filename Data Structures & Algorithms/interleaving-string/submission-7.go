func isInterleave(s1 string, s2 string, s3 string) bool {
	cache := make(map[[2]int]bool)
	m, n := len(s1), len(s2)
	if m + n != len(s3){
		return false
	}
	var isInterleaveFromIdx func(idx1, idx2 int)bool
	isInterleaveFromIdx = func(idx1, idx2 int)bool{
		if idx1 == m && idx2 == n {
			return true
		}
		key := [2]int{idx1,idx2}
		if val, existed := cache[key]; existed{
			return val
		}
		result := false
		idx3 := idx1 + idx2
		if idx1 < m && s1[idx1] == s3[idx3]{
			result = isInterleaveFromIdx(idx1+1,idx2)
		}
		if idx2 < n && s2[idx2] == s3[idx3]{
			result = isInterleaveFromIdx(idx1, idx2+1)
		}
		cache[key]=result
		return result
	}
	return isInterleaveFromIdx(0,0)
}
