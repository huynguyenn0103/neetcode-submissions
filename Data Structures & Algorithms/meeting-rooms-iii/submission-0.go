type OriginalItem struct{
	startTime int
	endTime int
}

type OriginalMinHeap []OriginalItem
func (h OriginalMinHeap) Len()int{return len(h)}
func (h OriginalMinHeap) Less(i, j int)bool {return h[i].startTime < h[j].startTime}
func (h OriginalMinHeap) Swap(i, j int){h[i],h[j]=h[j],h[i]}
func (h *OriginalMinHeap) Push(x interface{}){*h = append(*h, x.(OriginalItem))}
func (h *OriginalMinHeap) Pop()interface{}{
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}


type RealItem struct{
	startTime int
	endTime int
	room int
}
type RealMinHeap []RealItem
func (h RealMinHeap) Len()int{return len(h)}
func (h RealMinHeap) Less(i,j int)bool{
	if h[i].endTime == h[j].endTime{
		return h[i].room < h[j].room
	}
	return h[i].endTime < h[j].endTime
}
func (h RealMinHeap) Swap(i,j int){h[i],h[j]=h[j],h[i]}
func (h *RealMinHeap) Push(x interface{}){*h = append(*h, x.(RealItem))}
func (h *RealMinHeap) Pop()interface{}{
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

type FreeRoom struct {
	room int
}
type RoomHeap []FreeRoom
func (h RoomHeap) Len()int {return len(h)}
func (h RoomHeap) Less(i,j int)bool {return h[i].room < h[j].room}
func (h RoomHeap) Swap(i,j int){h[i],h[j]=h[j],h[i]}
func (h *RoomHeap) Push(x interface{}){*h = append(*h, x.(FreeRoom))}
func (h *RoomHeap) Pop()interface{}{
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}


func mostBooked(n int, meetings [][]int) int {
	roomUsed := make([]int,n)
	originalMinHeap := &OriginalMinHeap{}
	heap.Init(originalMinHeap)
	for _, meeting := range meetings{
		heap.Push(originalMinHeap, OriginalItem{startTime: meeting[0], endTime: meeting[1]})
	}
	roomMinHeap := &RoomHeap{}
	heap.Init(roomMinHeap)
	for i := 0; i < n; i++ {
		heap.Push(roomMinHeap, FreeRoom{room: i})
	}

	realMinHeap := &RealMinHeap{}
	heap.Init(realMinHeap)
	
	nextTime := (*originalMinHeap)[0].startTime 
	for originalMinHeap.Len() > 0{
		for realMinHeap.Len() > 0 && (*realMinHeap)[0].endTime == nextTime{
			item := heap.Pop(realMinHeap).(RealItem)
			heap.Push(roomMinHeap, FreeRoom{room: item.room})
		}
		for originalMinHeap.Len() > 0 && (*originalMinHeap)[0].startTime <= nextTime && roomMinHeap.Len() > 0{
			item := heap.Pop(originalMinHeap).(OriginalItem)
			roomItem := heap.Pop(roomMinHeap).(FreeRoom)
			roomUsed[roomItem.room]++
			heap.Push(realMinHeap, RealItem{startTime: nextTime, endTime: nextTime + item.endTime-item.startTime, room: roomItem.room})
		}
		if realMinHeap.Len() > 0 {
			nextTime = (*realMinHeap)[0].endTime
		}
		if originalMinHeap.Len() > 0 && roomMinHeap.Len() > 0{
			if realMinHeap.Len() == 0 {
				nextTime = (*originalMinHeap)[0].startTime
			} else {
				nextTime = min((*realMinHeap)[0].endTime, (*originalMinHeap)[0].startTime)
			}
		}
	}
	maxIdx := 0
	maxVal := 0
	for i, val := range roomUsed{
		if val > maxVal {
			maxIdx = i
			maxVal = val
		}
	}
	return maxIdx
}
