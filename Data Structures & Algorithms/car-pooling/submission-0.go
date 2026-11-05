type ActionType string
const (Drop ActionType = "DROP"
		Pick ActionType = "PICK")

type HeapItem struct {
	action ActionType
	location int
	numOfPassengers int
}

type MinHeap []HeapItem
func (h MinHeap) Len()int{return len(h)}
func (h MinHeap) Less(i,j int)bool{
	if h[i].location == h[j].location {
		if h[i].action == h[j].action || h[i].action == "DROP" && h[j].action == "PICK"{
			return true
		} else {
			return false
		}
	} else {
		return h[i].location < h[j].location
	}
}
func (h MinHeap) Swap(i,j int){h[i],h[j]=h[j],h[i]}
func (h *MinHeap) Push(x interface{}){*h = append(*h, x.(HeapItem))}
func (h *MinHeap) Pop()interface{}{
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}


func carPooling(trips [][]int, capacity int) bool {
	minHeap := &MinHeap{}
	for _, trip := range trips {
		heap.Push(minHeap, HeapItem{action: Pick, location: trip[1], numOfPassengers: trip[0]})
		heap.Push(minHeap, HeapItem{action: Drop, location: trip[2], numOfPassengers: trip[0]})
	}
	for minHeap.Len() > 0 {
		top := heap.Pop(minHeap).(HeapItem)
		if top.action == Pick {
			capacity = capacity - top.numOfPassengers
		} else {
			capacity = capacity + top.numOfPassengers
		}
		if capacity < 0 {return false}
	}
	return true
}
