import "slices"

func subsets(nums []int) [][]int {
	result := make([][]int,0)
	subset := make([]int,0)
	n := len(nums)
	var buildSubsetFromIdx func(idx int)
	buildSubsetFromIdx = func(idx int){
		if idx >= n {
			result = append(result, slices.Clone(subset))
			return
		}
		buildSubsetFromIdx(idx+1)
		subset = append(subset, nums[idx])
		buildSubsetFromIdx(idx+1)
		subset = subset[:len(subset)-1]
	}
	buildSubsetFromIdx(0)
	return result
}
