func rob(nums []int) int {
    cacheMaxMoneyFromIdx := make(map[int]int)
	n := len(nums)
	var maxMoneyFromIdx func(idx int)int
	maxMoneyFromIdx = func(idx int)int{
		if idx >= n {
			return 0
		}
		if val, existed := cacheMaxMoneyFromIdx[idx]; existed{
			return val
		}
		result := nums[idx] + maxMoneyFromIdx(idx+2)
		result = max(result, maxMoneyFromIdx(idx+1))
		cacheMaxMoneyFromIdx[idx]=result
		return result
	}
	return maxMoneyFromIdx(0)
}
