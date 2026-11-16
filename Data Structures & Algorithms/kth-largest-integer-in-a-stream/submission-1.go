
type MaxHeap []int
func (h MaxHeap) Len()int{return len(h)}
func (h MaxHeap) Less(i,j int)bool { return h[i] > h[j]}
func (h MaxHeap) Swap(i,j int){h[i],h[j]=h[j],h[i]}
func (h *MaxHeap) Push(x interface{}){*h = append(*h, x.(int))}
func (h *MaxHeap) Pop()interface{}{
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

type KthLargest struct {
    k int
	maxHeap *MaxHeap
}


func Constructor(k int, nums []int) KthLargest {
    maxHeap := &MaxHeap{}
	for _, num := range nums{
		heap.Push(maxHeap, num)
	}
	return KthLargest{k: k, maxHeap: maxHeap}
}


func (this *KthLargest) Add(val int) int {
    heap.Push(this.maxHeap, val)
	tmp := make([]int, 0)
	k := this.k
	for k > 0 && this.maxHeap.Len() > 0 {
		num := heap.Pop(this.maxHeap).(int)
		tmp = append(tmp, num)
		k--
	}
	for _, num := range tmp{
		heap.Push(this.maxHeap,num)
	}
	return tmp[len(tmp)-1]
}
