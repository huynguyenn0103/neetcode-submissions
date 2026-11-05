func rangeBitwiseAnd(left int, right int) int {
	result := left
	for i := left + 1; i <= right; i++ {
		result = result & i
	}
	return result
}
