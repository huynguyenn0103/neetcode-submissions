func subarraySum(nums []int, k int) int {
	var preFixMap = map[int]int{}
	preFixMap[0] = 1
	res := 0
	currSum := 0
	for _, n := range nums {
		currSum += n
		if count, ok := preFixMap[currSum-k]; ok {
			res += count
		}
		preFixMap[currSum]++
	}
	return res
}