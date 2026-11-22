func square(x int)int{
	return x * x
}

func squareOfDist(point []int)int{
	return square(point[0]) + square(point[1])
}

type MinHeap [][]int
func (h MinHeap) Len()int{return len(h)}
func (h MinHeap) Less(i,j int)bool{return squareOfDist(h[i]) < squareOfDist(h[j])}
func (h MinHeap) Swap(i,j int){h[i],h[j] = h[j],h[i]}
func (h *MinHeap) Push(x interface{}){*h = append(*h, x.([]int))}
func (h *MinHeap) Pop()interface{}{
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

func kClosest(points [][]int, k int) [][]int {
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	for _, point := range points{
		heap.Push(minHeap, point)
	}
	result := make([][]int,0)
	for k > 0 {
		point := heap.Pop(minHeap).([]int)
		result = append(result, point)
		k--
	}
	return result
}
