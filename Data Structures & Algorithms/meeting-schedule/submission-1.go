/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

import "slices"

func isOverlap(i1,i2 Interval)bool{
	return i1.end > i2.start
}

func canAttendMeetings(intervals []Interval) bool {
	slices.SortFunc(intervals, func(i1,i2 Interval)int{
		if i1.start == i2.start{
			return i1.end - i2.end
		}
		return i1.start - i2.start
	})
	n := len(intervals)
	for i := 1; i < n; i++{
		if isOverlap(intervals[i-1], intervals[i]){
			return false
		}
	}
	return true
}
