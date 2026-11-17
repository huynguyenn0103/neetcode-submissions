func lengthOfLIS(nums []int) int {
    cacheLISFromIdx := make(map[int]int)
	n := len(nums)
	var lisFromIdx func(idx int)int
	lisFromIdx = func(idx int)int{
		if idx >= n {return 0}
		if result, existed := cacheLISFromIdx[idx]; existed{
			return result
		}
		result := 1
		for i := idx + 1; i < n; i++ {
			if nums[i] > nums[idx]{
				result = max(result, 1 + lisFromIdx(i))
			}
		}
		cacheLISFromIdx[idx] = result
		return result
	}
	for i := 0; i < n; i++{
		lisFromIdx(i)
	}
	result := 0
	for _, val := range cacheLISFromIdx{
		result = max(result,val)
	}
	return result
}
