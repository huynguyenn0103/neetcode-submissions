func isMatch(s string, p string) bool {
    cache := make(map[[2]int]bool)
	m,n := len(s),len(p)
	var recursion func(idx1, idx2 int)bool
	recursion = func(idx1, idx2 int)bool {
		if idx1 >= m && idx2 >= n {
			return true
		} else if idx1 >= m || idx2 >= n{
			if idx2 < n-1 && p[idx2+1]=='*'{
				return recursion(idx1, idx2+2)
			}
			return false
		}
		key := [2]int{idx1,idx2}
		if val, existed := cache[key]; existed{
			return val
		}
		isSuccess := false
		if idx2 + 1 < n{
			if p[idx2+1]=='*'{
				if s[idx1] == p[idx2] || p[idx2] == '.'{
					isSuccess = isSuccess || recursion(idx1+1,idx2) || recursion(idx1+1,idx2+2) || recursion(idx1, idx2+2)
				} else {
					isSuccess = isSuccess || recursion(idx1, idx2+2)
				}
			} else {
				if s[idx1] == p[idx2] || p[idx2] == '.'{
					isSuccess = isSuccess || recursion(idx1+1,idx2+1)
				}
			}
		} else {
			if s[idx1] == p[idx2] || p[idx2] == '.'{
				isSuccess = isSuccess || recursion(idx1+1,idx2+1)
			}
		}
		cache[key]=isSuccess
		return cache[key]
	}
	return recursion(0,0)
}
