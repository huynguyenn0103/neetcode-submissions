func plusOne(digits []int) []int {
    n := len(digits)
	for i := 0; i < n/2; i++ {
		digits[i], digits[n-1-i] = digits[n-1-i], digits[i]
	}
	newVal := digits[0] + 1
	digits[0] = newVal % 10
	carry := newVal / 10
	for i := 1; i < n; i++ {
		newVal := digits[i] +carry
		digits[i] = newVal % 10
		carry = newVal / 10
	}
	if carry == 1{
		digits = append(digits, 1)
	}
	m := len(digits)
	for i := 0; i < m/2; i++ {
		digits[i], digits[m-1-i] = digits[m-1-i], digits[i]
	}
	return digits
}
