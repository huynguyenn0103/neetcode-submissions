func countBit(n int) int{
	count := 0
	for n != 0 {
		if n & 1 == 1 {
			count++
		}
		n = n >> 1
	}
	return count
}

func countBits(n int) []int {
	result := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		result[i] = countBit(i)
	}
	return result
}
