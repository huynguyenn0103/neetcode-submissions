func myPow(x float64, n int) float64 {
    needDivide := n < 0
	if n < 0 {
		n = -n
	}
	result := 1.0
	for n > 0 {
		if n % 2 == 1{
			result = result * x
		} 
		x = x * x
		n = n /2
	}
	if needDivide {
		return 1.0/result
	}
	return result
}
