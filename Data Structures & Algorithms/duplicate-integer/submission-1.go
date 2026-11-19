func hasDuplicate(nums []int) bool {
    mapNum := make(map[int]bool)
	for _, num := range nums{
		if mapNum[num]{
			return true
		}
		mapNum[num]=true
	}
	return false
}
