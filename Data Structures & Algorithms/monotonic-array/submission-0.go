func isMonotonic(nums []int) bool {
    n := len(nums)
	if n <= 1 {return true}
	isIncrease := false
	isDecrease := false
	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1]{continue}
		if nums[i] > nums[i-1]{
			isIncrease = true
		} else if nums[i] < nums[i-1]{
			isDecrease = true
		}
	}
	if isDecrease && isIncrease{
		return false
	}
	return true
}