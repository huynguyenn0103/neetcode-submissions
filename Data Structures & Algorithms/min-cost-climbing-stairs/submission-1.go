func minCostClimbingStairs(cost []int) int {
    n := len(cost)
	cacheMinCostFromIdx := make(map[int]int)
	var getMinCostFromIdx func(idx int)int
	getMinCostFromIdx = func(idx int)int{
		if idx >= n {
			return 0
		}
		if val, existed := cacheMinCostFromIdx[idx]; existed{
			return val
		}
		result := math.MaxInt
		result = min(result, cost[idx] + getMinCostFromIdx(idx+1))
		result = min(result, cost[idx] + getMinCostFromIdx(idx+2))
		cacheMinCostFromIdx[idx] = result
		return result
	}
	getMinCostFromIdx(0)
	return min(cacheMinCostFromIdx[0],cacheMinCostFromIdx[1])
}
