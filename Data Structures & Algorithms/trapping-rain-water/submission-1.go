func calculateWater(from, to int, height []int)int{
	var target int
	if height[from] <= height[to]{
		target = height[from]
	} else {
		target = height[to]
	}
	ptr := from + 1
	result := 0
	for ptr < to {
		result += target - height[ptr]
		ptr++
	}
	return result
}

func trap(height []int) int {
	result := 0
	n := len(height)
	left, right := 0, n-1
	for left < right {
		if height[left] <= height[right]{
			ptr := left + 1
			for ptr <= right && height[ptr] < height[left]{
				ptr++
			}
			if height[ptr] >= height[left]{
				result += calculateWater(left, ptr, height)
			}
			left = ptr
		} else {
			ptr := right-1
			for ptr >= left && height[ptr] < height[right]{
				ptr--
			}
			if height[ptr] >= height[right]{
				result += calculateWater(ptr, right, height)
			}
			right = ptr
		}
	}
	return result
}
