func findDisappearedNumbers(nums []int) []int {
	mapNum := make(map[int]bool)
	for _, num := range nums {
		mapNum[num]=true
	}
	n := len(nums)
	result := make([]int, 0)
	for i := 1; i <= n; i++ {
		if mapNum[i] == false{
			result = append(result,i)
		}
	}
	return result
}
