func canPartitionKSubsets(nums []int, k int) bool {
	n := len(nums)
	sum := 0
	for _, val := range nums{
		sum += val
	}
	if sum % k != 0 {return false}
	target := sum / k
	group := make([]int,k)

	var backtrack func(idx int)bool
	backtrack = func(idx int)bool {
		if idx == n {
			return true
		}
		for i := 0; i < k; i++ {
			if group[i] + nums[idx] > target {continue}
			group[i] = group[i] + nums[idx]
			if backtrack(idx+1){
				return true
			}
			group[i] = group[i] - nums[idx]
			if group[i] == 0 {break}
		}
		return false
	}
	return backtrack(0)
}
