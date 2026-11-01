type MinHeap []int
func (h MinHeap) Len()int {return len(h)}
func (h MinHeap) Less(i, j int)bool{return h[i]<h[j]}
func (h MinHeap) Swap(i, j int){h[i],h[j]=h[j],h[i]}
func (h *MinHeap) Push(x interface{}){*h = append(*h, x.(int))}
func (h *MinHeap) Pop()interface{}{
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

func sortColors(nums []int) {
	minHeap := &MinHeap{}
	heap.Init(minHeap)
    for _, color := range nums{
		heap.Push(minHeap,color)
	}
	for i:= range nums{
		item := heap.Pop(minHeap).(int)
		nums[i]=item
	}
}
