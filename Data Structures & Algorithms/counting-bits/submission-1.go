func numOfOnes(x int)int{
	num := 0
	for x != 0 {
		num += x & 1
		x = x >> 1
	}
	return num
}

func countBits(n int) []int {
	result := make([]int, 0)
	for i := 0; i <= n; i++ {
		result = append(result, numOfOnes(i))
	}
	return result
}
