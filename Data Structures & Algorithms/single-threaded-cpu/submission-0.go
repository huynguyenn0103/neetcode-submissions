import "slices"
import "cmp"

type HeapItem struct {
	idx int
	pTime int
}
type MinHeap []HeapItem
func (h MinHeap) Len()int{return len(h)}
func (h MinHeap) Less(i,j int)bool{
	if h[i].pTime == h[j].pTime{
		return h[i].idx < h[j].idx
	}
	return h[i].pTime < h[j].pTime
}
func (h MinHeap) Swap(i,j int){h[i],h[j]=h[j],h[i]}
func (h *MinHeap) Push(x interface{}){*h = append(*h, x.(HeapItem))}
func (h *MinHeap) Pop()interface{}{
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

func getOrder(tasks [][]int) []int {
	tasksWithIdx := make([][]int, len(tasks))
	for i := range tasks {
		tasksWithIdx[i]=[]int{tasks[i][0],tasks[i][1],i}
	}

	slices.SortFunc(tasksWithIdx, func(t1,t2 []int)int{
		return cmp.Compare(t1[0],t2[0])
	})
	startTime := tasksWithIdx[0][0]
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	ptr := 0
	result := make([]int, 0)
	for {
		for ptr < len(tasksWithIdx) && tasksWithIdx[ptr][0] <= startTime{
			heap.Push(minHeap,HeapItem{idx: tasksWithIdx[ptr][2], pTime: tasksWithIdx[ptr][1]})
			ptr++
		}
		if minHeap.Len() == 0 {break}
		item := heap.Pop(minHeap).(HeapItem)
		result = append(result,item.idx)
		startTime = startTime + item.pTime
		if ptr < len(tasksWithIdx){
			startTime = max(startTime, tasksWithIdx[ptr][0])
		}
	}
	return result
}
