func majorityElement(nums []int) int {
    mapOfFreq := make(map[int]int)
	for _, val := range nums{
		mapOfFreq[val]++
		if mapOfFreq[val] > len(nums)/2{
			return val
		}
	}
	return -1
}
