func findMin(nums []int) int {
	minVal := nums[0]
	n := len(nums)
	left, right := 0, n-1
	for left <= right{
		mid := (left + right)/2
		if nums[left] < nums[mid] && nums[mid] < nums[right] {
			return nums[left]
		} else if nums[mid] <= nums[left] && nums[mid] <= nums[right]{
			minVal = min(minVal, nums[mid])
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return minVal
}
