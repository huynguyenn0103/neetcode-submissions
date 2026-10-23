func squareOfDistToOrigin(p1 []int)int {
	return p1[0]*p1[0] + p1[1]*p1[1]
}

type MaxHeap [][]int
func (h MaxHeap) Len()int {return len(h)}
func (h MaxHeap) Less(i, j int)bool{return squareOfDistToOrigin(h[i]) < squareOfDistToOrigin(h[j])}
func (h MaxHeap) Swap(i, j int){h[i],h[j]=h[j],h[i]}
func (h *MaxHeap) Push(x interface{}){*h = append(*h, x.([]int))}
func (h *MaxHeap) Pop()interface{}{
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

func kClosest(points [][]int, k int) [][]int {
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	for _, point := range points{
		heap.Push(maxHeap, point)
	}
	result := make([][]int, 0)
	for k > 0 {
		result = append(result, heap.Pop(maxHeap).([]int))
		k--
	}
	return result
}
