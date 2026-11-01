func getConcatenation(nums []int) []int {
    n := len(nums)
	for i := 0; i < n; i++ {
		nums = append(nums, nums[i])
	}
	return nums
}
