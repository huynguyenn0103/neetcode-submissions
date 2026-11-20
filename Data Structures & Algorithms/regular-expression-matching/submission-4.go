func isMatch(s string, p string) bool {
    cacheResultFromIdx := make(map[[2]int]bool)
	m,n := len(s), len(p)
	var isMatchFromIdx func(idx1, idx2 int)bool
	isMatchFromIdx = func(idx1, idx2 int)bool{
		if idx1 >= m && idx2 >= n {
			return true
		}
		if idx2 >= n {
			return false
		}
		key := [2]int{idx1,idx2}
		if val, existed := cacheResultFromIdx[key]; existed{
			return val
		}
		result := false
		if idx1 >= m{
			if idx2+1 < n && p[idx2+1] == '*'{
				result = isMatchFromIdx(idx1,idx2+2)
			}
		} else {
			if p[idx2] == '.'{
				if idx2 + 1 < n && p[idx2+1] == '*'{
					result = result || isMatchFromIdx(idx1+1, idx2)
					result = result || isMatchFromIdx(idx1+1, idx2+2)
					result = result || isMatchFromIdx(idx1, idx2+2)
				} else {
					result = result || isMatchFromIdx(idx1+1, idx2+1)
				}
			} else {
				if s[idx1] == p[idx2]{
					if idx2 + 1 < n && p[idx2+1] == '*'{
						result = result || isMatchFromIdx(idx1+1,idx2)
						result = result || isMatchFromIdx(idx1+1,idx2+2)
						result = result || isMatchFromIdx(idx1,idx2+2)
					} else {
						result = result || isMatchFromIdx(idx1+1, idx2+1)
					}
				} else {
					if idx2 + 1 < n && p[idx2+1] == '*'{
						result = result || isMatchFromIdx(idx1, idx2+2)
					}
				}
			}
		}
		cacheResultFromIdx[key]=result
		return result
	}
	return isMatchFromIdx(0,0)
}
