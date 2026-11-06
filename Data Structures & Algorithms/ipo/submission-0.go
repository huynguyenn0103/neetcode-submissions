type HeapItem struct{
	capital int
	profit int
}

type MinHeap []HeapItem
func (h MinHeap) Len()int{return len(h)}
func (h MinHeap) Less(i,j int)bool{return h[i].capital < h[j].capital}
func (h MinHeap) Swap(i,j int){h[i],h[j]=h[j],h[i]}
func (h *MinHeap) Push(x interface{}){*h = append(*h,x.(HeapItem))}
func (h *MinHeap) Pop()interface{}{
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

type MaxHeap []HeapItem
func (h MaxHeap) Len()int{return len(h)}
func (h MaxHeap) Less(i,j int)bool{return h[i].profit > h[j].profit}
func (h MaxHeap) Swap(i,j int){h[i],h[j]=h[j],h[i]}
func (h *MaxHeap) Push(x interface{}){*h = append(*h,x.(HeapItem))}
func (h *MaxHeap) Pop()interface{}{
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	for i := 0; i < len(profits); i++ {
		heap.Push(minHeap, HeapItem{profit:profits[i],capital:capital[i]})
	}
	maxHeap := &MaxHeap{}
	for k > 0 {
		for minHeap.Len() > 0 {
			top := (*minHeap)[0]
			if top.capital > w {break}
			item := heap.Pop(minHeap).(HeapItem)
			heap.Push(maxHeap,item)
		}
		if maxHeap.Len() > 0 {
			item := heap.Pop(maxHeap).(HeapItem)
			w = w + item.profit
		}
		k--
	}
	return w
}
