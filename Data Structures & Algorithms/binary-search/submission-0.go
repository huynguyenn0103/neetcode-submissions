func search(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := (left + right)/2
		if nums[mid] == target{
			return mid
		} else if nums[mid] > target{
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
