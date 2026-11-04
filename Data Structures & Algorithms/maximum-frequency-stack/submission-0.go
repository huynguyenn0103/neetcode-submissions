type HeapItem struct {
	val int
	timestamp int
	freq int
}

type MaxHeap []HeapItem
func (h MaxHeap) Len()int{return len(h)}
func (h MaxHeap) Less(i,j int)bool {
	if h[i].freq == h[j].freq{
		return h[i].timestamp > h[j].timestamp
	}
	return h[i].freq > h[j].freq
}
func (h MaxHeap) Swap(i,j int){h[i],h[j]=h[j],h[i]}
func (h *MaxHeap) Push(x interface{}){*h = append(*h, x.(HeapItem))}
func (h *MaxHeap) Pop()interface{}{
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

var timestamp = 0
func getTS()int{
	timestamp++
	return timestamp
}

type FreqStack struct {
	mapValToFreq map[int]int
	maxHeap *MaxHeap
}

func Constructor() FreqStack {
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	return FreqStack{
		maxHeap: maxHeap,
		mapValToFreq: make(map[int]int),
	}
}

func (this *FreqStack) Push(val int) {
	this.mapValToFreq[val]++
	heap.Push(this.maxHeap, HeapItem{
		val: val,
		timestamp: getTS(),
		freq: this.mapValToFreq[val],
	})
}

func (this *FreqStack) Pop() int {
	top := heap.Pop(this.maxHeap).(HeapItem)
	this.mapValToFreq[top.val]--
	return top.val
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor()
 * obj.Push(val)
 * param2 := obj.Pop()
 */
 