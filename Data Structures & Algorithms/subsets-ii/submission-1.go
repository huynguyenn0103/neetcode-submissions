import "slices"
func isSame(r1, r2 []int)bool{
	if len(r1) != len(r2){return false}
	for i := 0; i < len(r1); i++{
		if r1[i] != r2[i]{return false}
	}
	return true
}

func getDistinctResult(result [][]int)[][]int{
	n := len(result)
	dResult := make([][]int,0)
	for i := 0; i < n; i++ {
		isAdded := true
		for j := i + 1; j < n; j++ {
			if isSame(result[i],result[j]){
				isAdded = false
				break
			}
		}
		if isAdded {
			dResult = append(dResult, result[i])
		}
	}
	return dResult
}

func subsetsWithDup(nums []int) [][]int {
	slices.Sort(nums)
	result := make([][]int,0)
	subsets := make([]int,0)
	n := len(nums)
	var buildSubsetFromIdx func(idx int)
	buildSubsetFromIdx = func(idx int){
		if idx >= n {
			result = append(result, slices.Clone(subsets))
			return
		}
		buildSubsetFromIdx(idx+1)
		subsets = append(subsets, nums[idx])
		buildSubsetFromIdx(idx+1)
		subsets = subsets[:len(subsets)-1]
	}
	buildSubsetFromIdx(0)
	return getDistinctResult(result)
}
