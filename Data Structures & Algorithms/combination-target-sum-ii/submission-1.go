import "slices"

func combinationSum2(candidates []int, target int) [][]int {
	n := len(candidates)
	slices.Sort(candidates)
	result := make([][]int,0)
	combination := make([]int,0)
	var buildTargetFromIdx func(idx int, target int)
	buildTargetFromIdx = func(idx int, target int){
		if target == 0 {
			result = append(result, slices.Clone(combination))
			return
		}
		if idx >= n || target < 0 {return}
		for i := idx; i < n; i++ {
			if i > idx && candidates[i] == candidates[i-1]{continue}
			combination = append(combination, candidates[i])
			buildTargetFromIdx(i+1, target - candidates[i])
			combination = combination[:len(combination)-1]
		}
	}
	buildTargetFromIdx(0,target)
	return result
}
