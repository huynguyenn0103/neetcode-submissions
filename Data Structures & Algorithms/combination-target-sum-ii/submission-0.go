import "slices"

// 1 2 2 4 5 6 9

func combinationSum2(candidates []int, target int) [][]int {
	slices.SortFunc(candidates, func(i1,i2 int)int{return i1-i2})
	result := make([][]int, 0)
	combination := make([]int, 0)
	
	var backtrack func(curIdx int, needSum int)
	backtrack = func(curIdx int, needSum int){
		if needSum == 0 {
			result = append(result, slices.Clone(combination))
		}
		if needSum < 0 {return}
		for i := curIdx; i < len(candidates); i++ {
			if i > curIdx && candidates[i] == candidates[i-1]{continue}
			combination = append(combination, candidates[i])
			backtrack(i+1, needSum-candidates[i])
			combination = combination[:len(combination)-1]
		}
	}
	backtrack(0, target)
	return result
}
