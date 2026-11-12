func divideArray(nums []int) bool {
    mapValToFreq := make(map[int]int)
	for _, val := range nums{
		mapValToFreq[val]++
	}
	for _, val := range mapValToFreq{
		if val % 2 == 1{return false}
	}
	return true
}