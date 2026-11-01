func subarraySum(nums []int, k int) int {
	var preFixMap = map[int]int{}
	preFixMap[0] = 1
	res := 0
	currSum := 0
	for _, n := range nums {
		currSum += n
		res += preFixMap[currSum-k]
		preFixMap[currSum]++
	}
	return res
}