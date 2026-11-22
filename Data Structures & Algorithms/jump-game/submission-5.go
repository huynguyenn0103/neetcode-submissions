func canJump(nums []int) bool {
    cache := make(map[int]bool)
	n := len(nums)
	var canJumpFromIdx func(idx int)bool
	canJumpFromIdx = func(idx int)bool{
		if idx >= n-1{
			return true
		}
		if val, existed := cache[idx]; existed{
			return val
		}
		result := false
		for i := idx+1; i <= idx+nums[idx]; i++{
			result = result || canJumpFromIdx(i)
		}
		cache[idx]=result
		return result
	}
	return canJumpFromIdx(0)
}
