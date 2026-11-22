func largestRectangleArea(heights []int) int {
    stack := make([]int, 0)
    result := 0

    for i := 0; i <= len(heights); i++ {
        curHeight := 0
        if i < len(heights) {
            curHeight = heights[i]
        }

        for len(stack) > 0 && heights[stack[len(stack)-1]] > curHeight {
            top := len(stack) - 1
            height := heights[stack[top]]
            stack = stack[:top]

            width := i
            if len(stack) > 0 {
                width = i - stack[len(stack)-1] - 1
            }

            result = max(result, height*width)
        }

        stack = append(stack, i)
    }

    return result
}