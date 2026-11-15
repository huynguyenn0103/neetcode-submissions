// 0 + 3 + 1 + 1 +

func XorAll(subset []int)int{
	if len(subset) > 0 {
		tmp := subset[0]
		for i := 1; i < len(subset);i++{
			tmp = tmp ^ subset[i]
		}
		return tmp
	} 
	return 0
}

func subsetXORSum(nums []int) int {
	result := 0
	subset := make([]int,0)
	n := len(nums)
	var backtrack func(idx int)
	backtrack = func(idx int){
		if idx >= n {
			result += XorAll(subset)
			return
		}
		subset = append(subset, nums[idx])
		backtrack(idx+1)
		subset = subset[:len(subset)-1]

		backtrack(idx+1)
	}
	backtrack(0)
	return result
}
