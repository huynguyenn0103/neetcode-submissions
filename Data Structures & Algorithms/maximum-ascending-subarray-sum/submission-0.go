func maxAscendingSum(nums []int) int {
    result := 0
	tmp := nums[0]
	for i := 1; i < len(nums); i++{
		if nums[i] <= nums[i-1]{
			result = max(result, tmp)
			tmp = nums[i]
		} else {
			tmp = tmp + nums[i]
		}
	}
	return max(tmp, result)
}