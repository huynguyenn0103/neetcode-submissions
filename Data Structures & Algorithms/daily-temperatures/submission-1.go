type StackItem struct{
	temp int
	idx int
}

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int,n)
	stack := make([]StackItem, 0)
	for idx, temp := range temperatures{
		for len(stack) > 0 && stack[len(stack)-1].temp < temp{
				item := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				result[item.idx] = idx - item.idx
			}
		stack = append(stack, StackItem{temp: temp, idx: idx})
	}
	return result
}
