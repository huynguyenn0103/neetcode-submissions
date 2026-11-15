import "slices"

func isSame(r1, r2 []int)bool{
	if len(r1) != len(r2) {return false}
	for i := 0; i < len(r1); i++{
		if r1[i] != r2[i]{
			return false
		}
	}
	return true
}

func getDistinctResult(result [][]int)[][]int{
	n := len(result)
	distinctResults := make([][]int,0)
	for i := 0; i < n; i++{
		same := false
		for j := i+1; j < n; j++{
			if isSame(result[i], result[j]){
				same = true
				break
			}
		}
		if !same{
			distinctResults = append(distinctResults, result[i])
		}
	}
	return distinctResults
}

func combinationSum(nums []int, target int) [][]int {
	n := len(nums)
    result := make([][]int,0)
	combination := make([]int,0)
	var combinationWithTarget func(idx int, target int)
	combinationWithTarget = func(idx int, target int){
		if target < 0 {return}
		if target == 0{
			clone := slices.Clone(combination)
			slices.Sort(clone)
			result = append(result, clone)
			return
		}
		if idx >= n {return}
		valAtIdx := nums[idx]
		maxAppearance := target / valAtIdx
		for i := 0; i <= maxAppearance; i++{
			if i == 0 {
				combinationWithTarget(idx+1,target)
			} else {
				combination = append(combination, valAtIdx)
				combinationWithTarget(idx+1, target - valAtIdx * i)
			}
		}
		combination = combination[:len(combination)-maxAppearance]
	}
	combinationWithTarget(0, target)
	return getDistinctResult(result)
}
