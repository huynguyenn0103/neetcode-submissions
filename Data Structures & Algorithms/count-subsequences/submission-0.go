// ccaaat  cat

func numDistinct(s string, t string) int {
    m, n := len(s), len(t)
	cache := make(map[[2]int]int)
	var recursion func(idx1, idx2 int)int
	recursion = func(idx1, idx2 int)int{
		if idx2 >= n {
			return 1
		}
		if idx1 >= m {
			return 0
		}
		key := [2]int{idx1,idx2}
		if val, existed := cache[key]; existed{
			return val
		}
		numOfWays := 0
		for i := idx1; i < m; i++ {
			if s[i] == t[idx2]{
				numOfWays += recursion(i+1,idx2+1)
			}
		}
		cache[key]=numOfWays
		return cache[key]
	}
	recursion(0,0)
	return cache[[2]int{0,0}]
}
