func minDistance(word1 string, word2 string) int {
    cacheMinOperationFromPos := make(map[[2]int]int)
	m, n := len(word1), len(word2)
	var minOperationFromPos func(idx1, idx2 int)int
	minOperationFromPos = func(idx1, idx2 int)int{
		if idx2 >= n {
			return max(0, m-idx1)
		}
		if idx1 >= m {
			return n - idx2
		}
		key := [2]int{idx1,idx2}
		if result, existed := cacheMinOperationFromPos[key]; existed{
			return result
		}
		minResult := math.MaxInt
		if word1[idx1] == word2[idx2]{
			minResult = minOperationFromPos(idx1+1,idx2+1)
		}
		minResult = min(minResult, 1 + minOperationFromPos(idx1+1, idx2+1))
		minResult = min(minResult, 1 + minOperationFromPos(idx1+1, idx2))
		minResult = min(minResult, 1 + minOperationFromPos(idx1, idx2+1))
		cacheMinOperationFromPos[key]=minResult
		return minResult
	}
	return minOperationFromPos(0,0)
}
