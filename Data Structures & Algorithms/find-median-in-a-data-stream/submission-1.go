type MinHeap []int
func (h MinHeap) Len()int{return len(h)}
func (h MinHeap) Less(i, j int)bool{return h[i] < h[j]}
func (h MinHeap) Swap(i, j int){h[i],h[j]=h[j],h[i]}
func (h* MinHeap) Push(x interface{}){*h = append(*h, x.(int))}
func (h *MinHeap) Pop()interface{}{
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

type MaxHeap []int
func (h MaxHeap) Len()int{return len(h)}
func (h MaxHeap) Less(i, j int)bool{return h[i] > h[j]}
func (h MaxHeap) Swap(i, j int){h[i],h[j]=h[j],h[i]}
func (h* MaxHeap) Push(x interface{}){*h = append(*h, x.(int))}
func (h *MaxHeap) Pop()interface{}{
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

const Infinity = math.MaxInt

type MedianFinder struct {
    minHeap *MinHeap
	maxHeap *MaxHeap
}


func Constructor() MedianFinder {
    minHeap := &MinHeap{}
	maxHeap := &MaxHeap{}
	heap.Init(minHeap)
	return MedianFinder{minHeap: minHeap, maxHeap: maxHeap}
}

// mean 2
// maxHeap 1 
// minHeap 3 2

func (this *MedianFinder) AddNum(num int)  {
    var topMaxHeap int
	if this.maxHeap.Len() == 0 {
		topMaxHeap = Infinity
	} else {
		topMaxHeap = (*this.maxHeap)[0]
	}
	
	if num < topMaxHeap {
		heap.Push(this.maxHeap, num)
	} else {
		heap.Push(this.minHeap, num)
	}
	if this.maxHeap.Len() - this.minHeap.Len() == 2{
		heap.Push(this.minHeap, heap.Pop(this.maxHeap).(int))
	}
	if this.minHeap.Len() - this.maxHeap.Len() == 2{
		heap.Push(this.maxHeap, heap.Pop(this.minHeap).(int))
	}
}


func (this *MedianFinder) FindMedian() float64 {
    n, m := this.minHeap.Len(), this.maxHeap.Len()
	if n == 0 {
		return float64((*this.maxHeap)[0])
	}
	topMaxHeap := (*this.maxHeap)[0]
	topMinHeap := (*this.minHeap)[0]
	if n == m {
		return (float64(topMaxHeap) + float64(topMinHeap))/2
	} else if n > m {
		return float64(topMinHeap)
	} else {
		return float64(topMaxHeap)
	}
}
