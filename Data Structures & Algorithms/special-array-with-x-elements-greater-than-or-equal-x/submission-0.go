// (x1,...xn) >= n

func specialArray(nums []int) int {
    n := len(nums)
	for i := 0; i <= n; i++ {
		count := 0
		for _, num := range nums{
			if num >= i {
				count++
			}
		}
		if count == i{
			return count
		}
	}
	return -1

}