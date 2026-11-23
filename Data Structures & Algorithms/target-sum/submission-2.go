func findTargetSumWays(nums []int, target int) int {
    cache := make(map[[2]int]int)
	n := len(nums)
	var numOfWayFromTarget func(idx int, target int)int
	numOfWayFromTarget = func(idx int, target int)int{
		if idx >= n {
			if target == 0 {
				return 1
			}
			return 0
		}
		key := [2]int{idx,target}
		if val, existed := cache[key]; existed{
			return val
		}
		result := numOfWayFromTarget(idx+1,target-nums[idx])+numOfWayFromTarget(idx+1,target+nums[idx])
		cache[key]=result
		return result
	}
	return numOfWayFromTarget(0,target)
}
