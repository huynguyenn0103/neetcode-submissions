func rangeBitwiseAnd(left int, right int) int {
	steps := 0
	for left!=right{
		left = left>>1
		right = right>>1
		steps++
	}

	return left<<steps
}