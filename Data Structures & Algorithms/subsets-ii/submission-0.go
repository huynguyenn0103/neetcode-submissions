// 1 1 2
import "slices"

func subsetsWithDup(nums []int) [][]int {
	slices.SortFunc(nums, func(i1,i2 int)int{return i1-i2})
	result := make([][]int, 0)
	subset := make([]int, 0)
	var backtrack func(curIdx int)
	backtrack = func(curIdx int){
		result = append(result, slices.Clone(subset))
		for i := curIdx; i < len(nums); i++ {
			if i > curIdx && nums[i] == nums[i-1]{continue}
			subset = append(subset, nums[i])
			backtrack(i+1)
			subset = subset[:len(subset)-1]
		}
	}
	backtrack(0)
	return result
}
