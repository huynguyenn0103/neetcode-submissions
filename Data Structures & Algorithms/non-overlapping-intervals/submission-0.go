import "slices"

func isOverlapAndGetCandidate(i1, i2 []int) ([]int, bool){
	i1To := i1[1]
	i2From, i2To := i2[0], i2[1]
	if  i1To > i2From {
		if i1To >= i2To {
			return i2, true
		}
		return i1, true
	}
	return i2, false
}

func eraseOverlapIntervals(intervals [][]int) int {
    // (1, 7) (1, 8) (2 3)
	slices.SortFunc(intervals, func(i1, i2 []int)int{
		if i1[0] == i2[0]{
			return i1[1] - i2[1]
		}
		return i1[0] - i2[0]
	})
	candidateItv := intervals[0]
	n := len(intervals)
	count := 0
	for i := 1; i < n; i++ {
		var isOverlap bool
		candidateItv, isOverlap = isOverlapAndGetCandidate(candidateItv, intervals[i])
		if isOverlap{
			count++
		}
	}
	return count
}
