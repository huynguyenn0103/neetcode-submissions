type MinHeap []int
func (h MinHeap) Len()int{return len(h)}
func (h MinHeap) Less(i, j int)bool{return h[i] < h[j]}
func (h MinHeap) Swap(i, j int){h[i],h[j]=h[j],h[i]}
func (h *MinHeap) Push(x interface{}){*h = append(*h, x.(int))}
func (h *MinHeap) Pop()interface{}{
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}


func leastInterval(tasks []byte, n int) int {
	stepRequired := n + 1
	mapTaskToLatestTime := make(map[byte]int)
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	for _, task := range tasks {
		if mapTaskToLatestTime[task] == 0 {
			mapTaskToLatestTime[task] = 1
		} else {
			mapTaskToLatestTime[task] = mapTaskToLatestTime[task] + stepRequired
		}
		heap.Push(minHeap, mapTaskToLatestTime[task])
	}
	result := 0
	// (1),(1),(2)
	for minHeap.Len() != 0 {
		top := heap.Pop(minHeap).(int)
		if top > result {
			result = top
		} else {
			result++
		}
	}
	return result
}
