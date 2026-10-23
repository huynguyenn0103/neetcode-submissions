/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */
import "slices"

type MinHeap []Interval
func (h MinHeap) Len() int{return len(h)}
func (h MinHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h MinHeap) Less(i, j int)bool{return h[i].end < h[j].end}
func (h *MinHeap) Push(x interface{}){*h = append(*h, x.(Interval))}
func (h *MinHeap) Pop()interface{}{
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

func isOverlap(i1, i2 Interval) bool{
	toI1 := i1.end
	fromI2 := i2.start
	if toI1 <= fromI2 {
		return false
	}
	return true
}

func minMeetingRooms(intervals []Interval) int {
	if len(intervals) == 0 {
		return 0
	}
	slices.SortFunc(intervals, func(i1, i2 Interval)int {
		if i1.start == i2.start{
			return i1.end - i2.end
		}
		return i1.start - i2.start
	})
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	heap.Push(minHeap, intervals[0])
	n := len(intervals)
	for i := 1; i < n; i++ {
		top := (*minHeap)[0]
		if isOverlap(top, intervals[i]){
			heap.Push(minHeap, intervals[i])
		} else {
			heap.Pop(minHeap)
			heap.Push(minHeap, intervals[i])
		}
	}
	return minHeap.Len()
}
