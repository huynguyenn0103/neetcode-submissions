import "slices"
func permute(nums []int) [][]int {
	visitedIdx := make(map[int]bool)
	result := make([][]int, 0)
	combination := make([]int, 0)
	var backtrack func()
	backtrack = func(){
		if len(combination) == len(nums){
			result = append(result, slices.Clone(combination))
		}
		for i := 0; i < len(nums); i++ {
			if visitedIdx[i] {continue}
			visitedIdx[i]=true
			combination = append(combination, nums[i])
			backtrack()
			visitedIdx[i]=false
			combination = combination[:len(combination)-1]
		}
	}
	backtrack()
	return result
}
