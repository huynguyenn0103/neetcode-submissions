func twoSum(nums []int, target int) []int {
    mapValToIdx := make(map[int]int)
	for i, num := range nums{
        if idx, existed := mapValToIdx[target-num]; existed{
			return []int{idx,i}
		}
		mapValToIdx[num]= i
	}
	return []int{}
}
