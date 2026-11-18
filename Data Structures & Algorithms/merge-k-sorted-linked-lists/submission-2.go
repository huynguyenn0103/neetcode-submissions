/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// k * n
// log(n) * n

type MinHeap []*ListNode
func (h MinHeap) Len()int{return len(h)}
func (h MinHeap) Less(i,j int)bool{return h[i].Val < h[j].Val}
func (h MinHeap) Swap(i,j int){h[i],h[j] = h[j],h[i]}
func (h *MinHeap) Push(x interface{}){*h = append(*h, x.(*ListNode))}
func (h *MinHeap) Pop()interface{}{
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {

    minHeap := &MinHeap{}
	heap.Init(minHeap)
	for _, node := range lists{
		if node != nil {
			heap.Push(minHeap, node)
		}
	}
	var result *ListNode
	cur := result
	for minHeap.Len()>0 {
		node := heap.Pop(minHeap).(*ListNode)
		newNode := &ListNode{Val: node.Val}
		if result == nil {
			result = newNode
			cur = result
		} else {
			cur.Next = newNode
			cur = cur.Next
		}
		node = node.Next
		if node != nil{
			heap.Push(minHeap, node)
		}
	}
	return result
}
