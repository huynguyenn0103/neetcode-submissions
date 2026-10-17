func findMin(nums []int) int {
	n := len(nums)
	left, right := 0, n-1
	result := nums[0]
	top := nums[0]
	for left <= right{
		mid := (left + right)/2
		if nums[mid] >= top {
			left = mid + 1
		} else if nums[mid] < top {
			right = mid - 1
		}
		result = min(result, nums[mid])
	}
	return result
}
