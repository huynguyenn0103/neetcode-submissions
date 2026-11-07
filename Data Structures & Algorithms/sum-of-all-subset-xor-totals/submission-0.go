func subsetXORSum(nums []int) int {
	result := 0
	combination := 0
	n := len(nums)
	var xorAllSubsetFromI func(idx int)
	xorAllSubsetFromI = func(idx int){
		if idx >= n {
			result += combination
			return
		}
		result += combination
		for i := idx; i < n; i++ {
			tmp := combination
			combination = combination ^ nums[i]
			xorAllSubsetFromI(i+1)
			combination = tmp
		}
	}
	xorAllSubsetFromI(0)
	return result
}
