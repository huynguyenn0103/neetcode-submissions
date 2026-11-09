
// 12 

func integerBreak(n int) int {
	cache := make(map[int]int)
	var maxProductFromNum func(num int)int
	maxProductFromNum = func(num int)int{
		if num <= 0 {
			return 1
		}
		if val, existed := cache[num]; existed{
			return val
		}
		var result int
		if num == n {
			result = 0
		} else {
			result = num
		}

		for i := 1; i < num; i++ {
			result = max(result, i*maxProductFromNum(num-i))
		}
		cache[num]=result
		return cache[num]
	}
	return maxProductFromNum(n)
}
