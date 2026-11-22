func trap(height []int) int {
	result := 0
	n := len(height)
	stack := make([]int,0)
	for i := 0; i < n; i++{
		if len(stack) == 0 {
			stack = append(stack,i)
		} else {
			target := min(height[stack[0]], height[i])
			for len(stack) > 0 && height[i] >= height[stack[len(stack)-1]]{
				l := len(stack)
				topIdx := stack[l-1]
				stack = stack[:l-1]
				if len(stack) > 0 {
					beforeTopIdx := stack[len(stack)-1]
					result += (target-height[topIdx])*(topIdx-beforeTopIdx)
				}
			}
			stack = append(stack,i)
		}
	}
	return result
}
