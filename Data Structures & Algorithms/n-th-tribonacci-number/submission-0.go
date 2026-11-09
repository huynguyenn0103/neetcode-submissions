func tribonacci(n int) int {
	cache := make(map[int]int)
	var calculate func(n int)int
	calculate = func(n int)int{
		if n == 0 {return 0}
		if n <= 2 {
			return 1
		}
		if val, isExisted := cache[n]; isExisted{
			return val
		}
		result := calculate(n-1) + calculate(n-2) + calculate(n-3)
		cache[n]=result
		return cache[n]
	}
	return calculate(n)
}
