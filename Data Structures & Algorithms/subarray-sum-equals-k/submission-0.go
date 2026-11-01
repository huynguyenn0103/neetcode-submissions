
func subarraySum(nums []int, k int) int {
	n := len(nums)
	result := 0
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j <n; j++ {
			sum = sum + nums[j]
			if sum == k{
				result++
			}
		}
	}
	return result
}
