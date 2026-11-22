func largestRectangleArea(heights []int) int {
	maxHeight := -math.MaxInt
	for _, height := range heights{
		maxHeight = max(maxHeight,height)
	}
	n := len(heights)
	result := 0
	for i := 1; i <= maxHeight; i++ {
		left := -1
		for right:=0; right < n; right++{
			if heights[right] > 0 {
				if left == -1{
					left = right
				}
				heights[right]--
			} else {
				if left != -1{
					result = max(result, (right-left)*i)
					left = -1
				}
			}
		}
		if left != -1{
			result = max(result, (n-left)*i)
		}
	}
	return result
}
