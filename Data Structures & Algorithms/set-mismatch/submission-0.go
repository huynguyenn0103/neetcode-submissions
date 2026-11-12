func findErrorNums(nums []int) []int {
	var misNum, dupNum int
    mapChar := make(map[int]bool)
	n := len(nums)
	for i := 0; i < n; i++ {
		if mapChar[nums[i]]{
			dupNum = nums[i]
		} else {
			mapChar[nums[i]]=true
		}
	}
	for i := 1; i <= n; i++ {
		if !mapChar[i]{
			misNum = i
			break
		}
	}
	return []int{dupNum, misNum}
}