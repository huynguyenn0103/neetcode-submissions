func reverseBits(n int) int {
	result := 0
	for i := 1; i <= 32; i++{
		last := n & 1
		n = n >> 1
		result = result << 1 | last
	}
	return result
}
