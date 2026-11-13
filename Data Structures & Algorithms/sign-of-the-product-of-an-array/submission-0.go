func signFunc(x int)int{
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	} else {
		return 0
	}
}

func arraySign(nums []int) int {
    result := 1
	for _, num := range nums{
		result = result * signFunc(num)
	}
	return result
}