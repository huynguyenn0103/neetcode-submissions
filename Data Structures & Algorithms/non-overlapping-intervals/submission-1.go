import "slices"

func isOverlap(i1, i2 []int)bool{
	return i2[0] < i1[1]
}

func eraseOverlapIntervals(intervals [][]int) int {
	slices.SortFunc(intervals, func(i1,i2 []int)int{
		if i1[0] == i2[0]{
			return i1[1] - i2[1]
		}
		return i1[0] - i2[0]
	})
    removed := 0
	var latestInterval []int
	for _, interval := range intervals{
		if latestInterval != nil && isOverlap(latestInterval, interval) {
			removed++
			if latestInterval[1] > interval[1]{
				latestInterval = interval
			}
		} else {
			latestInterval = interval
		}
	}
	return removed
}
