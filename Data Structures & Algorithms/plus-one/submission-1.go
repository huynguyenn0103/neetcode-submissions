func plusOne(digits []int) []int {
    result := make([]int,0)
	n := len(digits)
	carry := (digits[n-1] + 1) / 10
	result = append(result, (digits[n-1] + 1) % 10)
	for i := n-2; i >= 0; i--{
		sum := (digits[i] + carry) % 10
		carry = (digits[i] + carry) / 10
		result = append(result, sum)
	}
	if carry == 1{
		result = append(result,1)
	}
	for i := 0; i < len(result)/2; i++{
		result[i],result[len(result)-1-i]=result[len(result)-1-i],result[i]
	}
	return result
}
