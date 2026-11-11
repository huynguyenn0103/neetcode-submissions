func longestMonotonicSubarray(nums []int) int {
	result := 1
    isIncrease := false
	n := len(nums)
	tmp := 1
	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1]{
			result = max(result, tmp)
			tmp = 1
		} else if (nums[i] > nums[i-1] && isIncrease) || (nums[i-1] > nums[i] && !isIncrease){
			tmp++
		} else {
			if nums[i] > nums[i-1]{
				isIncrease = true
			} else {
				isIncrease = false
			}
			result = max(result, tmp)
			tmp = 2
		}
	}
	return max(result, tmp)
}