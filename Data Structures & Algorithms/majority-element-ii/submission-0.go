func majorityElement(nums []int) []int {
	mapToFreq := make(map[int]int)
	for _, val := range nums{
		mapToFreq[val]++
	}
	n := len(nums)
	result := make([]int,0)
	for val, freq := range mapToFreq{
		if freq > n/3{
			result = append(result, val)
		} 
	}
	return result
}
