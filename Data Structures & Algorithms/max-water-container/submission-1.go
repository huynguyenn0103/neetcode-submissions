func maxArea(heights []int) int {
	n := len(heights)
	left, right := 0, n-1
	result := 0
	for left <= right {
		result = max(result, (right-left)*min(heights[left],heights[right]))
		if heights[left] == heights[right]{
			left++
			right--
		} else if heights[left] > heights[right]{
			right--
		} else {
			left++
		}
	}
	return result
}
