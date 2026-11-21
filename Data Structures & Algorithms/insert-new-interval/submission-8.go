func isBetween(i1, i2, target []int) bool{
	return target[0] > i1[1] && target[1] < i2[0]
}

func isOverlap(i1, i2 []int) bool{
	return !(i1[1] < i2[0] || i2[1] < i1[0])
}

func insert(intervals [][]int, newInterval []int) [][]int {
    result := make([][]int,0)
	n := len(intervals)
	if n == 0 {
		return [][]int{newInterval}
	}
	if isBetween([]int{-1,-1}, intervals[0], newInterval){
		result = append(result, newInterval)
	}

	for i := 0; i < n; i++ {
		if i >= 1 && isBetween(intervals[i-1],intervals[i], newInterval) {
			result = append(result, newInterval)
		}
		start := intervals[i][0]
		end := intervals[i][1]
		ptr := i
		for ptr < n && isOverlap(intervals[ptr], newInterval){
			start = min(start, intervals[ptr][0], newInterval[0])
			end = max(end, intervals[ptr][1], newInterval[1])
			ptr++
		}
		result = append(result, []int{start,end})
		if ptr != i && ptr != n {
			i = ptr - 1
		} else {
			i = ptr
		}
	}

	if isBetween(intervals[n-1], []int{math.MaxInt,math.MaxInt}, newInterval){
		result = append(result, newInterval)
	}

	return result
}
