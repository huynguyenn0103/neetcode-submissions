func combinationSum4(nums []int, target int) int {
	cache := make(map[int]int)
	var combinationFrom func(num int)int
	combinationFrom = func(num int)int{
		if num == 0 {
			return 1
		}
		if num < 0 {
			return 0
		}
		if val, existed := cache[num]; existed{
			return val
		}

		result := 0
		for	_, val := range nums{
			result += combinationFrom(num-val)
		}
		cache[num]=result
		return result
	}
	return combinationFrom(target)
}
