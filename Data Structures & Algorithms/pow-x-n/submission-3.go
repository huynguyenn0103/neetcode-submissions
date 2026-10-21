func myPow(x float64, n int) float64 {
	if x == 1{
		return 1
	} 
	if x == -1 {
		if n % 2 == 0 {
			return 1
		}
		return -1
	}
	total := 1.0
	if n >= 0 {
		for n > 0 {
			total = total * x
			n--
		}
	} else {
		if n < -1000 {
			return 0
		}
		for n < 0 {
			total = total / x
			n++
		}
	}
	return total
}
