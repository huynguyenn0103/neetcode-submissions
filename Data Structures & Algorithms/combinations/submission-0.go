import "slices"
func combine(n int, k int) [][]int {
	result := make([][]int,0)
	combination := make([]int,0)
	var backtrack func(idx int)
	backtrack = func(idx int){
		if len(combination) == k {
			result = append(result, slices.Clone(combination))
			return
		}
		if idx > n {
			return
		}
		for i := idx; i <=n; i++ {
			combination = append(combination, i)
			backtrack(i+1)
			combination = combination[:len(combination)-1]
		}
	}
	backtrack(1)
	return result
}
