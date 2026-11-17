import "slices"

func isSame(r1, r2 []int)bool{
	for i := 0; i < len(r1); i++ {
		if r1[i] != r2[i]{return false}
	}
	return true
}

func getDistinctResult (results [][]int)[][]int{
	dResult := make([][]int,0)
	n := len(results)
	for i := 0; i < n; i++ {
		isAdded := true
		for j := i+1; j < n; j++ {
			if isSame(results[i],results[j]){
				isAdded = false
				break
			}
		}
		if isAdded{
			dResult = append(dResult, results[i])
		}
	}
	return dResult
}

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	n := len(nums)
	results := make([][]int, 0)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1]{continue}
		target := - nums[i]
		mapValToIdx := make(map[int]int)
		for j := i+1; j < n; j++ {
			if _, existed := mapValToIdx[target - nums[j]]; existed{
				results = append(results, []int{nums[i], target - nums[j], nums[j]})
			}
			mapValToIdx[nums[j]] = j
		}
	}
	return getDistinctResult(results)
}
