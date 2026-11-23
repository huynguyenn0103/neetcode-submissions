func longestConsecutive(nums []int) int {
    cache := make(map[int]int)
	mapNum := make(map[int]bool)
	for _, num := range nums{
		mapNum[num]=true
	}
	var longestFromNum func(num int) int
	longestFromNum = func(num int) int{
		if mapNum[num]==false{
			return 0
		}
		if val, existed := cache[num]; existed{
			return val
		}
		result := 1 + longestFromNum(num+1)
		cache[num]=result
		return result
	}
	result := 0
	for _, num := range nums{
        result = max(result, longestFromNum(num))
	}
	return result
}
