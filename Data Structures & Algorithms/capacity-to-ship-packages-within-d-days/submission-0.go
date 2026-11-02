// 9,10 => 1 1 1 2


func shipWithinDays(weights []int, days int) int {
	n := len(weights)
	dayNeededWithShipCap := func (capacity int) int {
		days := 1
		ptr := 0
		remainCap := capacity
		for ptr < n {
			if remainCap >= weights[ptr]{
				remainCap = remainCap - weights[ptr]
			} else {
				days++
				remainCap = capacity
				if remainCap >= weights[ptr]{
					remainCap = remainCap - weights[ptr]
				}
			}
			ptr++
		}
		return days
	}

	leftCap, rightCap := 0, 0
	for _, val := range weights {
		rightCap += val
		leftCap = max(leftCap, val)
	}
	for leftCap <= rightCap {
		mid := (leftCap + rightCap)/2
		day := dayNeededWithShipCap(mid)
		if day > days {
			leftCap = mid + 1
		} else {
			rightCap = mid - 1
		}
	}
	return leftCap
}
