type Item struct {
	val int
	idx int
}

func dailyTemperatures(temperatures []int) []int {
	stack := make([]Item, 0)
	result := make([]int, len(temperatures))
	for i, val := range temperatures{
		for len(stack) > 0 && stack[len(stack)-1].val < val{
			result[stack[len(stack)-1].idx]= i - stack[len(stack)-1].idx
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, Item{val: val, idx: i})
	}
	return result
}
