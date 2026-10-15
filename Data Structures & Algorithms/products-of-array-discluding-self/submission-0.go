func productExceptSelf(nums []int) []int {
	n := len(nums)
	afterIMul := make([]int, n-1)

	// [48, 24, 6]
	tmp := 1
	for i := n-1; i >= 1; i--{
		afterIMul[i-1]= nums[i] * tmp
		tmp = afterIMul[i-1]
	}

	result := make([]int, n)

	beforeIMul := 1
	for i, val := range nums {
		if i == n-1 {
			result[i] = beforeIMul
		} else {
			result[i] = beforeIMul * afterIMul[i]
		}
		beforeIMul = beforeIMul * val
	}
	return result
}
