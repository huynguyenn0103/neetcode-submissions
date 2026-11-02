// -3,-3,0,0,1,1,2,2,3,3
import "slices"
func fourSum(nums []int, target int) [][]int {
	cache := make(map[string]bool)
	slices.SortFunc(nums, func(i1,i2 int)int{return i1-i2})
	n := len(nums)
	result := make([][]int, 0)
	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			sum := target - nums[i] - nums[j]
			left, right := j+1,n-1
			for left < right {
				curSum := nums[left]+nums[right]
				if curSum > sum {
					right--
				} else if curSum < sum {
					left++
				} else {
					key := fmt.Sprintf("%v:%v:%v:%v",nums[i],nums[j],nums[left],nums[right])
					if !cache[key]{
						result = append(result, []int{nums[i],nums[j],nums[left],nums[right]})
						cache[key]=true
					}
					left++
					right--
				}
			}
		}
	}
	return result
}
