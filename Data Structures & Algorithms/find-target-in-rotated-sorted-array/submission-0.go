func search(nums []int, target int) int {
	n := len(nums)
	left, right := 1, n-1
	top := nums[0]
	if target == top{return 0}
	for left <= right{
		mid := (left + right)/2
		if target == nums[mid]{
			return mid
		}
		if nums[mid] < top {
			if target < nums[mid]{
				right = mid - 1
			} else if target < top {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if target > nums[mid]{
				left = mid + 1
			} else if target > top {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}

//     top    mid
//target => left = mid + 1
//        target => right = mid-1
//					target => left = mid + 1