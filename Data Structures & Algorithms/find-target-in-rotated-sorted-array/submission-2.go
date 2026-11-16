func search(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right{
		mid := (left + right)/2
		if nums[mid] == target{
			return mid
		}
		if nums[mid] > target {
			if nums[mid] >= nums[left]{
				if target == nums[left]{
					return left
				} else if target > nums[left]{
					right = mid - 1
				} else {
					left = mid + 1
				}
			} else {
				right = mid - 1
			}
		} else {
			if nums[mid] >= nums[left]{
				left = mid + 1
			} else {
				if target == nums[left]{
					return left
				} else if target > nums[left]{
					right = mid - 1
				} else {
					left = mid + 1
				}
			}
		}
	}
	return -1
}
