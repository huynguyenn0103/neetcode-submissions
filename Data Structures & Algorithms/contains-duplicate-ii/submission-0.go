func containsNearbyDuplicate(nums []int, k int) bool {
	numAndLatestPosition := make(map[int]int)
	for i, val := range nums{
		if pos, existed := numAndLatestPosition[val]; existed {
			if i - pos <= k {
				return true
			}
		} 
		numAndLatestPosition[val]=i
	}
	return false
}
