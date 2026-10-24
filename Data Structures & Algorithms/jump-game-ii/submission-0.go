func jump(nums []int) int {
	n := len(nums)
	if n == 1{
		return 0
	}
    result := 1
	start, end := 0, nums[0]
	if end >= n-1{
		return result
	}
	for {
		next := 0
		result++
		for start <= end {
			next = max(next, start + nums[start])
			if next >= n-1 {
				return result
			}		
			start++
		}
		end = next
	}
	return -1
}
