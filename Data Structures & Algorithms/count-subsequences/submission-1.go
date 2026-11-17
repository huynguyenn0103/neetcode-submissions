func numDistinct(s string, t string) int {
    cacheNumOfWayFromIdx := make(map[[2]int]int)
	m,n := len(s),len(t)
	var numOfWayFromIdx func(idx1, idx2 int)int
	numOfWayFromIdx = func(idx1, idx2 int)int{
		if idx2 >= n{
			return 1
		}
		if idx1 >= m {
			return 0
		}
		key := [2]int{idx1,idx2}
		if result, existed := cacheNumOfWayFromIdx[key]; existed{
			return result
		}
		result := 0
		if s[idx1] == t[idx2]{
			result += numOfWayFromIdx(idx1+1,idx2+1)
		}
		result += numOfWayFromIdx(idx1+1,idx2)
		cacheNumOfWayFromIdx[key]=result
		return result
	}
	return numOfWayFromIdx(0,0)
}
