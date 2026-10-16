func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	left, right := 0, n-1
	for left < right {
		if numbers[right] + numbers[left] < target{
			left++
		} else if numbers[right] + numbers[left] > target{
			right--
		} else {
			return []int{left+1, right+1}
		}
	}
	return []int{}
}
