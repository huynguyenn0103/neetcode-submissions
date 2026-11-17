func maxSubArray(nums []int) int {
    result := nums[0]
	n := len(nums)
	sum := nums[0]
	for i := 1; i < n; i++{
		result = max(result, sum)
		sum = sum + nums[i]
		if sum < nums[i]{
			sum = nums[i]
		}
	}
	return max(result, sum)
}
