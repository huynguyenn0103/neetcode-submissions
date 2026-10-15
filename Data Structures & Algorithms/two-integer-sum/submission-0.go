func twoSum(nums []int, target int) []int {
    mapValueToIndex := make(map[int]int)
	for i, val := range nums {
		if idx, existed := mapValueToIndex[target-val]; existed {
			return []int{idx, i}
		} else {
			mapValueToIndex[val]=i
		}
	}
	return []int{}
}
