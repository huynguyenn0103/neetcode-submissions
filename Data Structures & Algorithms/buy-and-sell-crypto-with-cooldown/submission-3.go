func maxProfit(prices []int) int {
    maxProfitFromIdxCache := make(map[int]int)
	n := len(prices)
	var maxProfitFromIdx func(idx int)int
	maxProfitFromIdx = func(idx int)int{
		if idx >= n {return 0}
		if result, existed := maxProfitFromIdxCache[idx]; existed{
			return result
		}
		result := 0
		for i:=idx+1; i<n; i++{
			if prices[i] > prices[idx]{
				result = max(result, prices[i] - prices[idx] + maxProfitFromIdx(i+2))
			}
		}
		result = max(result, maxProfitFromIdx(idx+1))
		maxProfitFromIdxCache[idx]=result
		return result
	}
	return maxProfitFromIdx(0)
}
