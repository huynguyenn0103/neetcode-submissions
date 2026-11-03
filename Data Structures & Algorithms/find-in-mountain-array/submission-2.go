/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, mountainArr *MountainArray) int {
    n := mountainArr.length()
	var turnPoint int
	left, right := 0, n-1
	for left <= right {
		mid := (left + right)/2
		if mid == 0 {
			left = mid + 1
		} else if mid == n-1 {
			right = mid - 1
		} else {
			valBefore := mountainArr.get(mid-1)
			valAfter := mountainArr.get(mid+1)
			valMid := mountainArr.get(mid)
			if valBefore > valMid && valMid > valAfter {
				right = mid - 1
			} else if valBefore < valMid && valMid < valAfter{
				left = mid + 1
			} else if valBefore < valMid && valMid > valAfter{
				turnPoint = mid
				break
			}
		}
	}
	left, right = 0, turnPoint
	for left <= right {
		mid := (left + right)/2
		valMid := mountainArr.get(mid)
		if valMid == target {
			return mid
		} else if valMid > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	left, right = turnPoint, n-1
	for left <= right {
		mid := (left + right)/2
		valMid := mountainArr.get(mid)
		if valMid == target {
			return mid
		} else if valMid < target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
