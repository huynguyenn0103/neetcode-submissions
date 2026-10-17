type Item struct {
	pos int
	val int
}

type MaxHeap []Item
func (h MaxHeap) Len()int {return len(h)}
func (h MaxHeap) Less(i, j int)bool {return h[i].val > h[j].val}
func (h MaxHeap) Swap(i, j int){h[i],h[j]=h[j],h[i]}
func (h *MaxHeap) Push(x interface{}){*h = append(*h, x.(Item))}
func (h *MaxHeap) Pop()interface{}{
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

func maxSlidingWindow(nums []int, k int) []int {
	result := make([]int, 0)
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
    ptr1, ptr2 := 0, k-1
	for i := ptr1; i < ptr2; i++ {
		heap.Push(maxHeap, Item{pos: i, val: nums[i]})
	}
	for ptr2 < len(nums) {
		heap.Push(maxHeap, Item{pos: ptr2, val: nums[ptr2]})
		var topItem Item
		for {
			topItem = ((*maxHeap)[0])
			if topItem.pos >= ptr1{break}
			heap.Pop(maxHeap)
		}
		result = append(result, topItem.val)
		ptr1++
		ptr2++
	}
	return result
}
