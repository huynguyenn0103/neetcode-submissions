func jump(nums []int) int {
    cache := make(map[int]int)
	n := len(nums)
	var minFromIdx func(idx int)int
	minFromIdx = func(idx int)int{
		if idx >= n-1 {
			return 0
		}
		if val, existed := cache[idx]; existed{
			return val
		}
		result := math.MaxInt
		for i := idx + 1; i <= idx + nums[idx]; i++{
			minVal := minFromIdx(i)
			if minVal == math.MaxInt{continue}
			result = min(result, 1 + minVal)
		}
		cache[idx]=result
		return result
	}
	return minFromIdx(0)
}
