/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

import "slices"

func canAttendMeetings(intervals []Interval) bool {
	slices.SortFunc(intervals, func(i1, i2 Interval)int {
		if i1.start == i2.start{
			return i1.end - i2.end
		}
		return i1.start - i2.start
	})
	n := len(intervals)
	for i := 1; i < n; i++ {
		if intervals[i].start < intervals[i-1].end{
			return false
		}
	}
	return true
}
