func singleNumber(nums []int) int {
	result := nums[0]
	n := len(nums)
	for i := 1; i < n; i++ {
		result = ^(result^nums[i])
	}
	return result
}
