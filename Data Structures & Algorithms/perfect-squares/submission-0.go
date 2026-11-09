func sqrt(x int)int{
	return int(math.Sqrt(float64(x)))
}

func isSquare(x int)bool {
	a := sqrt(x)
	return a * a == x
}

func numSquares(n int) int {
	cache := make(map[int]int)
	var leastSquareToSum func(num int)int
	leastSquareToSum = func(num int)int{
		if isSquare(num){
			return 1
		}
		if val, isExisted := cache[num]; isExisted{
			return val
		}
		x := sqrt(num)
		result := math.MaxInt
		for i := x; i >= 1; i-- {
			result = min(result, 1 + leastSquareToSum(num - i*i))
		}
		cache[num]=result
		return result
	}
	return leastSquareToSum(n)
}
