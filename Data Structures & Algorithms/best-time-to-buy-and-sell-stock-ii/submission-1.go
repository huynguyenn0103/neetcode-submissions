func maxProfit(prices []int) int {
	n := len(prices)
	cache := make(map[int]int)
	var maxProfitBuyFromIdx func(idx int)int
	maxProfitBuyFromIdx = func(idx int)int{
		if idx >= n {return 0}
		if val, existed := cache[idx]; existed{
			return val
		}
		maxVal := 0
		for i := idx; i < n; i++ {
			for j := i + 1; j < n; j++ {
				if prices[j] > prices[i]{
					maxVal = max(maxVal, prices[j]-prices[i] + maxProfitBuyFromIdx(j+1))
				}
			}
		}
		cache[idx]=maxVal
		return cache[idx]
	}
	for i := range prices{
		maxProfitBuyFromIdx(i)
	}
	tmp := 0
	for _, val := range cache {
		tmp = max(tmp, val)
	}
	return tmp
}
