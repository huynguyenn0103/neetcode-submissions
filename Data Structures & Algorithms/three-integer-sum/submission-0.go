import "slices"

func threeSum(nums []int) [][]int {
	slices.SortFunc(nums, func(num1, num2 int)int{return num1-num2})
	n := len(nums)
	result := make([][]int, 0)
	for i := 0; i < n; i++{
		if i >= 1 && nums[i]==nums[i-1]{continue}
		target := 0 - nums[i]
		left, right := i+1, n-1
		for left < right{
			if nums[left] + nums[right] > target{
				right--
			} else if nums[left] + nums[right] < target{
				left++
			} else {
				result = append(result, []int{nums[left], nums[right], nums[i]})
				left++
				right--
				for left > 0 && left < n && nums[left] == nums[left-1]{
					left++
				}
				for right >= 0 && right < n-1 && nums[right] == nums[right+1]{
					right--
				}
			}
		}
	}
	return result
}
