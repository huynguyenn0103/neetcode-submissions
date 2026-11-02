func removeDuplicates(nums []int) int {
	result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if i >= 1 && nums[i]==nums[i-1]{continue}
		result = append(result, nums[i])
	}
	for i := range result {
		nums[i]=result[i]
	}
	return len(result)
}
