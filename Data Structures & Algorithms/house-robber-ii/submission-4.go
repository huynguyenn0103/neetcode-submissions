func maxFrom(nums []int) int{
	n := len(nums)
    cacheMaxMoneyFromIdx := make(map[int]int)
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
	maxMoneyFromIdx(0)
	return cacheMaxMoneyFromIdx[0]
}


func rob(nums []int) int {
	if len(nums) == 1{
		return nums[0]
	}
	max1 := maxFrom(nums[1:])
	max2 := maxFrom(nums[:len(nums)-1])
	return max(max1, max2)
}
