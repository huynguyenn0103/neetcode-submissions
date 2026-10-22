func reverseBits(n int) int {
	result := 0
	numOfBits := 32
	for numOfBits > 0 {
		lastBit := n & 1
		result = (result << 1) | lastBit
		n = n >> 1
		numOfBits--
	}
	return result
}
