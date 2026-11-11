func findMaxConsecutiveOnes(nums []int) int {
	result := 0
	tmp := 0
	for _, num := range nums {
		if num == 1 {
			tmp++
		} else {
			result = max(result, tmp)
			tmp = 0
		}
	}
	return max(result,tmp)
}
