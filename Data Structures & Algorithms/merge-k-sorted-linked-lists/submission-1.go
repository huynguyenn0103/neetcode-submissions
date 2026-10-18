/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


type MinHeap []*ListNode
func (h MinHeap) Len() int{return len(h)}
func (h MinHeap) Less(i, j int)bool {
	return h[i].Val < h[j].Val
}
func (h MinHeap) Swap(i, j int) {h[i],h[j]=h[j],h[i]}
func (h *MinHeap) Push(x interface{}){*h = append(*h, x.(*ListNode))}
func (h *MinHeap) Pop() interface{}{
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}


func mergeKLists(lists []*ListNode) *ListNode {
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	for _, head := range lists {
		if head == nil{continue}
		heap.Push(minHeap, head)
	}

	var result, rCur *ListNode

	for minHeap.Len() > 0 {
		item := heap.Pop(minHeap).(*ListNode)
		if result == nil {
			result = item
			rCur = item
		} else {
			rCur.Next = item
			rCur = rCur.Next
		}
		if item.Next != nil {
			heap.Push(minHeap, item.Next)
		}
	}
	return result
}
