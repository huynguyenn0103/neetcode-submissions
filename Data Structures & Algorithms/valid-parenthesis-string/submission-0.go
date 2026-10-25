type StackItem struct{
	p byte
	idx int
}

func checkValidString(s string) bool {
    starIndices := make([]int, 0)
	stack := make([]StackItem, 0)
	n := len(s)
	for i := 0; i < n ; i++ {
		c := s[i]
		if c == '*'{
			starIndices = append(starIndices, i)
			continue
		}
		if len(stack) > 0 {
			top := stack[len(stack)-1]
			if top.p == '(' && c == ')'{
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, StackItem{p: c, idx: i})
			}
		} else {
			stack = append(stack, StackItem{p: c, idx: i})
		}
	}
	for {
		if len(stack) == 0 {return true}
		if len(starIndices) == 0 {return false}
		top := stack[len(stack)-1]
		tailStarIdx := starIndices[len(starIndices)-1]
		if top.p == '('{
			if top.idx > tailStarIdx {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			if top.idx > tailStarIdx {
				stack = stack[:len(stack)-1]
			}
		}
		starIndices = starIndices[:len(starIndices)-1]
	}
	return len(stack) == 0
}

// *)))**