const Infinity = math.MaxInt

func abs(a,b int)int{
	if a > b {return a-b}
	return b-a
}

func findClosestElements(arr []int, k int, x int) []int {
	left, right := Infinity, Infinity
	n := len(arr)
	for i, val := range arr {
		if val >= x {
			right = i
			break
		}
	}
	if right == Infinity {
		left = n-1
	} else {
		left = right-1
	}

	resultRight := make([]int, 0)
	resultLeft := make([]int, 0)
	for left >= 0 && right <= n-1 && k > 0 {
		distLeft := abs(x, arr[left])
		distRight := abs(x, arr[right])
		if distLeft <= distRight {
			resultLeft = append(resultLeft, arr[left])
			left--
		} else {
			resultRight = append(resultRight, arr[right])
			right++
		}
		k--
	}
	for left >= 0 && k > 0 {
		resultLeft = append(resultLeft, arr[left])
		left--
		k--
	}
	for right <= n-1 && k > 0 {
		resultRight = append(resultRight, arr[right])
		right++
		k--
	}
	result := make([]int, 0)
	for i := len(resultLeft)-1; i >= 0; i-- {
		result = append(result, resultLeft[i])
	}
	result = append(result, resultRight...)
	return result
}
