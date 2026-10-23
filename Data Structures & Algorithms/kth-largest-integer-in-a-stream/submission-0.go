type MaxHeap []int
func (h MaxHeap) Len()int{return len(h)}
func (h MaxHeap) Less(i, j int)bool{return h[i] > h[j]}
func (h MaxHeap) Swap(i, j int){h[i],h[j]=h[j],h[i]}
func (h *MaxHeap) Push(x interface{}){*h = append(*h, x.(int))}
func (h *MaxHeap) Pop()interface{}{
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

type KthLargest struct {
   maxHeap *MaxHeap
   k int  
}


func Constructor(k int, nums []int) KthLargest {
    maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	for i := 0; i < len(nums); i++ {
		heap.Push(maxHeap, nums[i])
	}
	return KthLargest{
		maxHeap: maxHeap,
		k: k,
	}
}


func (this *KthLargest) Add(val int) int {
    heap.Push(this.maxHeap, val)
	i := this.k - 1
	popValue := make([]int, 0)
	for i > 0 {
		i--
		popValue = append(popValue, heap.Pop(this.maxHeap).(int))
	}
	item := (*this.maxHeap)[0]
	for _, val := range popValue{
		heap.Push(this.maxHeap, val)
	}
	return item
}
