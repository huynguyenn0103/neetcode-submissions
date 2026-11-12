func check(nums []int) bool {
    n := len(nums)
	if n <= 1{return false}
	numOfBreaks := 0
	for i := 1; i < n; i++{
		if nums[i] < nums[i-1]{
			numOfBreaks++
			if numOfBreaks >= 2{
				return false
			}
		}
	}
	if numOfBreaks == 1{
		return nums[0] >= nums[n-1]
	}
	return true
}