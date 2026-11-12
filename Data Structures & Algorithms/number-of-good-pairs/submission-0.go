func numIdenticalPairs(nums []int) int {
    mapVal := make(map[int]int)
	result := 0
	for _, num := range nums{
		mapVal[num]++
	}
	for _, val := range mapVal{
		result += val * (val - 1) / 2
	}
	return result
}