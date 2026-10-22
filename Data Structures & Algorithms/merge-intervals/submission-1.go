import "slices"
func merge(intervals [][]int) [][]int {
    slices.SortFunc(intervals, func (i1,i2 []int)int {
		if i1[0]==i2[0]{
			return i1[1]-i2[1]
		}
		return i1[0]-i2[0]
	})
	result := make([][]int, 0)
	n := len(intervals)
	for i := 0; i < n; i++ {
		curStart, curEnd := intervals[i][0], intervals[i][1]
		from, to := curStart, curEnd
		for i <= n - 1{
			curStart, curEnd := intervals[i][0], intervals[i][1]
			if to < curStart {
				i--
				break
			}
			to = max(to, curEnd)
			i++
		}
		result = append(result, []int{from,to})
	}
	return result
}
