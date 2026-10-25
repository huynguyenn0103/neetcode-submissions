// 1 -> nil, 2, 3
// 2 -> nil, 3
// 3 -> nil

// bt(0,[]) -> i=0, bt(1,[1]) -> i=1 bt(2, [1,2]) ->i=2 bt(3, [1,2,3])
import "slices"
func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	curNums := make([]int, 0)
	var backtrack func(curIdx int)
	backtrack = func(curIdx int){
		result = append(result, slices.Clone(curNums))
		for i := curIdx; i < len(nums);i++{
			curNums = append(curNums, nums[i])
			backtrack(i+1)
			curNums = curNums[:len(curNums)-1]
		}
	}
	backtrack(0)
	return result
}
