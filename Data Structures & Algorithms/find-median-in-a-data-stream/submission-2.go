// maxHeap: 1 1 2 2 1
// minHeap: 4 5 9

type MaxHeap []int
func (h MaxHeap) Len()int {return len(h)}
func (h MaxHeap) Less(i, j int)bool{return h[i] > h[j]}
func (h MaxHeap) Swap(i, j int){h[i],h[j]=h[j],h[i]}
func (h *MaxHeap) Push(x interface{}){*h = append(*h, x.(int))}
func (h *MaxHeap) Pop()interface{}{
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

type MinHeap []int
func (h MinHeap) Len()int {return len(h)}
func (h MinHeap) Less(i, j int)bool{return h[i] < h[j]}
func (h MinHeap) Swap(i, j int){h[i],h[j]=h[j],h[i]}
func (h *MinHeap) Push(x interface{}){*h = append(*h, x.(int))}
func (h *MinHeap) Pop()interface{}{
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}


type MedianFinder struct {
    minHeap *MinHeap
	maxHeap *MaxHeap
}


func Constructor() MedianFinder {
    minHeap := &MinHeap{}
	maxHeap := &MaxHeap{}
	heap.Init(minHeap)
	heap.Init(maxHeap)
	return MedianFinder{
		minHeap: minHeap,
		maxHeap: maxHeap,
	}
}


func (this *MedianFinder) AddNum(num int)  {
    heap.Push(this.maxHeap, num)
	if this.maxHeap.Len() - this.minHeap.Len() == 2{
		num := heap.Pop(this.maxHeap).(int)
		heap.Push(this.minHeap, num)
	}
	if this.maxHeap.Len() > 0 && this.minHeap.Len() > 0 {
		topMax := (*this.maxHeap)[0]
		topMin := (*this.minHeap)[0]
		if topMax > topMin {
			heap.Pop(this.maxHeap)
			heap.Pop(this.minHeap)
			heap.Push(this.maxHeap, topMin)
			heap.Push(this.minHeap, topMax)
		}
	}
}


func (this *MedianFinder) FindMedian() float64 {
    n1, n2 := this.maxHeap.Len(), this.minHeap.Len()
	if (n1 + n2) % 2 == 1{
		return float64((*this.maxHeap)[0])
	} else {
		return (float64((*this.maxHeap)[0]) + float64((*this.minHeap)[0]))/2.0
	}
}
