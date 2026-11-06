
func maxSubarraySumCircular(nums []int) int {
	if len(nums) == 1 {return nums[0]}
	n := len(nums)
	result := nums[0]
	ptr1 := 0
	for ptr1 < n{
		sum := 0
		ptr2 := ptr1
		for {
			sum = sum + nums[ptr2]
			result = max(result, sum)
			if sum < nums[ptr2]{
				break
			}
			ptr2 = (ptr2+1)%n
			if ptr2 == ptr1{
				break
			}
		}
		ptr1++
	}
	return result
}
