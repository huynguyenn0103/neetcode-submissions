import "slices"


func comparePermutations(nums1, nums2 []int)bool {
	n := len(nums1)
	for i := 0; i < n; i++ {
		if nums1[i] != nums2[i]{
			return false
		}
	}
	return true
}

func permuteUnique(nums []int) [][]int {
	slices.SortFunc(nums, func(i1,i2 int)int{return i1-i2})
	visitedIdx := make(map[int]bool)
	results := make([][]int,0)
	permutation := make([]int,0)
	n := len(nums)
	var backtrack func()
	backtrack = func(){
		if len(permutation) == n {
			clonePermutation := slices.Clone(permutation)
			results = append(results, clonePermutation)
		}
		for i := 0; i < n; i++ {
			if visitedIdx[i]{continue}
			visitedIdx[i]=true
			permutation = append(permutation, nums[i])
			backtrack()
			permutation = permutation[:len(permutation)-1]
			visitedIdx[i]=false
		}
	}
	backtrack()
	finalResult := make([][]int, 0)

	for i := 0; i < len(results); i++ {
		isDuplicate := false
		for _, val := range finalResult {
			isDuplicate = isDuplicate || comparePermutations(results[i],val)
		}
		if !isDuplicate{
			finalResult = append(finalResult, results[i])
		}
	}
	return finalResult
}
