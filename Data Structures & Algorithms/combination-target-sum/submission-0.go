// 3 -> [3,4,5]
// 4 -> [4, 5]
// 5 -> [5]

import "slices"

func combinationSum(nums []int, target int) [][]int {
    result := make([][]int, 0)
	combination := make([]int, 0)
	
	var backtrack func(curIdx int, needSum int)
	backtrack = func(curIdx int, needSum int){
		if needSum == 0 {
			result = append(result, slices.Clone(combination))
		}
		if needSum < 0 {return}
		for i := curIdx; i < len(nums); i++ {
			combination = append(combination, nums[i])
			backtrack(i, needSum-nums[i])
			combination = combination[:len(combination)-1]
		}
	}
	backtrack(0, target)
	return result
}
