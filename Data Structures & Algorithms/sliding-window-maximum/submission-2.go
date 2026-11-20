type Item struct{
	idx int
	val int
}
type MaxHeap []Item
func (h MaxHeap) Len()int {return len(h)}
func (h MaxHeap) Less(i, j int)bool{return h[i].val > h[j].val}
func (h MaxHeap) Swap(i, j int){h[i],h[j] = h[j],h[i]}
func (h *MaxHeap) Push(x interface{}){*h = append(*h, x.(Item))}
func (h *MaxHeap) Pop()interface{}{
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
    maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	for right := 0; right < k; right++{
		heap.Push(maxHeap, Item{idx: right, val: nums[right]})
	}
	result := make([]int,0)
	result = append(result, (*maxHeap)[0].val)
	for right := k; right < n; right++{
		heap.Push(maxHeap, Item{idx: right, val : nums[right]})
		left := right + 1 - k
		for maxHeap.Len() > 0 {
			topItem := (*maxHeap)[0]
			if topItem.idx < left{
				heap.Pop(maxHeap)
			} else {break}
		}
		result = append(result, (*maxHeap)[0].val)
	}
	return result
}
