func maxArea(heights []int) int {
	n := len(heights)
	left, right := 0, n-1
	result := 0
	for left < right {
		if heights[right] > heights[left]{
			result = max(result, heights[left] * (right-left))
			left++
		} else if heights[right] < heights[left]{
			result = max(result, heights[right] * (right-left))
			right--
		} else {
			result = max(result, heights[right] * (right-left))
			left++
			right--
		}
	}
	return result
}
