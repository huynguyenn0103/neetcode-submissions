func getSum(a int, b int) int {
	result := 0
	carry := 0
	count := 0
	for (a != 0 || b != 0) && count < 64 {
		tmp := 0
		lastBitA := a & 1
		lastBitB := b & 1
		if lastBitA == 1 && lastBitB == 1{
			tmp = 1
		}
		sum := lastBitA ^ lastBitB
		if sum == 1 && carry == 1{
			tmp = 1
		}
		sum = sum ^ carry
		result = result | (sum << count)
		carry = tmp
		count++
		a = a >> 1
		b = b >> 1
	}
	if carry == 1 && count != 64{
		result = result | (carry << count)
	}
	return result
}
