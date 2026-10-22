func insert(intervals [][]int, newInterval []int) [][]int {
    results := make([][]int, 0)
	n := len(intervals)
	newStart, newEnd := newInterval[0], newInterval[1]
	// (1,5),(7,9)
	// (4,6) | (4,8) | (6, 6.5) | (6,8)
	isCompleted := false
	for i := 0; i < n; i++ {
		curStart, curEnd := intervals[i][0], intervals[i][1]
		if newStart > curEnd || isCompleted {
			results = append(results, []int{curStart, curEnd})
		} else {
			from, to := -1, -1
			if newStart <= curStart {
				from = newStart
			} else {
				from = curStart
			}
			for {
				curStart, curEnd = intervals[i][0], intervals[i][1]
				if newEnd < curStart {
					i--
					to = newEnd
					break
				}
				if newEnd <= curEnd {
					to = curEnd
					break
				}
				i++
				if i == n {
					to = newEnd
					break
				}
			}
			results = append(results, []int{from, to})
			isCompleted = true
		}
	}
	if !isCompleted {
		results = append(results, []int{newStart, newEnd})
	}	
	return results
}
