type MinHeap []int
func (h MinHeap) Len()int{return len(h)}
func (h MinHeap) Swap(i,j int){h[i],h[j]=h[j],h[i]}
func (h MinHeap) Less(i, j int)bool{return h[i]<h[j]}
func (h *MinHeap) Push(x interface{}){*h = append(*h, x.(int))}
func (h *MinHeap) Pop()interface{}{
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}
func sortArray(nums []int) []int {
    minHeap := &MinHeap{}
	heap.Init(minHeap)
	result := make([]int, 0)
	for _, val := range nums{
		heap.Push(minHeap, val)
	}
	for minHeap.Len() > 0 {
		result = append(result, heap.Pop(minHeap).(int))
	}
	return result
}
