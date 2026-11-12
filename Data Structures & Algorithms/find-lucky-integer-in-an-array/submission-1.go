func findLucky(arr []int) int {
	mapValToFreq := make(map[int]int)
	for _, val := range arr{
		mapValToFreq[val]++
	}
	result := -math.MaxInt
	for key, val := range mapValToFreq{
		if val == key{
			result = max(result, val)
		}
	}
	if result == -math.MaxInt{
		return -1
	}
	return result
}
