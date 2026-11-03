func calPoints(operations []string) int {
	stack := make([]int, 0)
	for _, operation := range operations{
		if operation == "+"{
			num1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			num2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack,num2,num1,num1+num2)
		} else if operation == "D"{
			num := stack[len(stack)-1]
			stack = append(stack, num*2)
		} else if operation == "C"{
			stack = stack[:len(stack)-1]
		} else {
			num, _ := strconv.Atoi(operation)
			stack = append(stack, num)
		}
	}
	sum := 0
	for _, val := range stack {
		sum += val
	}
	return sum
}
