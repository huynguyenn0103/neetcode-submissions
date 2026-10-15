func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, token := range tokens{
		num, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, num)
		} else {
			n := len(stack)
			o2, o1 := stack[n-1], stack[n-2]
			stack = stack[:n-2]
			var tmp int
			if token == "+"{
				tmp = o1 + o2
			} else if token == "-"{
				tmp = o1 - o2
			} else if token == "*"{
				tmp = o1 * o2
			} else if token == "/"{
				tmp = o1 / o2
			}
			stack = append(stack, tmp)
		}
	}
	return stack[0]
}
