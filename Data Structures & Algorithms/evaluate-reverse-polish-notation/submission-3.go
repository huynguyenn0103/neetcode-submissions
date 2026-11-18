func toNum(str string)int{
	num, _ := strconv.Atoi(str)
	return num
}

func evalRPN(tokens []string) int {
	stack := make([]int,0)
	for _, token := range tokens{
		if token == "+" || token == "-" || token == "*" || token == "/"{
			num1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			num2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			var result int
			if token == "+"{
				result = num2 + num1
			} else if token == "-"{
				result = num2 - num1
			} else if token == "*"{
				result = num2 * num1
			} else {
				result = num2 / num1
			}
			stack = append(stack, result)
		} else {
			stack = append(stack, toNum(token))
		}
	}
	return stack[0]
}
