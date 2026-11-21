func canPartition(nums []int) bool {
    sum := 0
	for _, num := range nums{
		sum += num
	}
	if sum % 2 != 0 {return false}
	target := sum/2
	n := len(nums)

	cache := make(map[[2]int]bool)

	var sumToTarget func(idx int, target int)bool

	sumToTarget = func(idx, target int)bool{
		if target == 0 {
			return true
		}
		if idx >= n || target < 0{
			return false
		}
		key := [2]int{idx,target}
		if val, existed := cache[key]; existed{
			return val
		}
		result := false
		result = result || sumToTarget(idx+1, target)
		result = result || sumToTarget(idx+1, target - nums[idx])
		cache[key]=result
		return result
	}
	return sumToTarget(0,target)
}
