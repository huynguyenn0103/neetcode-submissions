type MaxHeap []int
func (h MaxHeap) Len()int{return len(h)}
func (h MaxHeap) Less(i, j int)bool{return h[i] > h[j]}
func (h MaxHeap) Swap(i,j int){h[i],h[j]=h[j],h[i]}
func (h *MaxHeap) Push(x interface{}){
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop()interface{}{
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

func lastStoneWeight(stones []int) int {
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	for _, stone := range stones{
		heap.Push(maxHeap, stone)
	}
	for maxHeap.Len() > 1 {
		topStone := heap.Pop(maxHeap).(int)
		secondStone := heap.Pop(maxHeap).(int)
		if secondStone < topStone{
			heap.Push(maxHeap, topStone-secondStone)
		}
	}
	if maxHeap.Len() == 0 {
		return 0
	}
	return (*maxHeap)[0]
}
