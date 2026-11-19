/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

import "slices"

type MinHeap []Interval
func (h MinHeap) Len()int{return len(h)}
func (h MinHeap) Less(i, j int)bool{return h[i].end < h[j].end}
func (h MinHeap) Swap(i, j int){h[i],h[j]=h[j],h[i]}
func (h *MinHeap) Push(x interface{}){*h = append(*h, x.(Interval))}
func (h *MinHeap) Pop()interface{}{
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

func minMeetingRooms(intervals []Interval) int {
	slices.SortFunc(intervals, func(i1, i2 Interval)int{
		return i1.start - i2.start
	})
	minHeap := &MinHeap{}
	requiredRoom := 0
	for _, interval := range intervals{
		for minHeap.Len() > 0 {
			top := (*minHeap)[0]
			if top.end <= interval.start{
				heap.Pop(minHeap)
			} else {
				break
			}
		}
		heap.Push(minHeap, interval)
		requiredRoom = max(requiredRoom, minHeap.Len())
	}
	return max(requiredRoom, minHeap.Len())
}
