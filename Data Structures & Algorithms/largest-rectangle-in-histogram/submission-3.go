type Item struct {
	position int
	height int
	origin int
}


func largestRectangleArea(heights []int) int {
	result := 0
	stack := make([]Item, 0)
	for i, height := range heights {
		target := Item{position: i, height: height, origin: i}
		for len(stack) > 0 {
			topItem := stack[len(stack)-1]
			if topItem.height <= target.height {
				if topItem.height == target.height{
					stack = stack[:len(stack)-1]
					target.position = topItem.position
				} else {
					target.position = topItem.origin + 1
				}
				break
			} else {
				result = max(result, topItem.height * (target.position-topItem.position))
				stack = stack[:len(stack)-1]
			}
			if len(stack)==0{
				target.position = 0
			}
		}
		stack = append(stack, target)
	}
	for _, item := range stack {
		result = max(result, (len(heights)-item.position) * item.height)
	}
	return result
}
