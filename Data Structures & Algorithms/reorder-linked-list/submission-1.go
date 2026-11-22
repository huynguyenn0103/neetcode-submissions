/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    mapIdxToNode := make(map[int]*ListNode)
	idx := 0
	cur := head
	for cur != nil{
		mapIdxToNode[idx]=cur
		idx++
		cur = cur.Next
	}
	left, right := 1, idx-1
	turn := 1
	cur = head
	for left <= right{
        if turn == 0{
			cur.Next = mapIdxToNode[left]
		   left++
		} else {
			cur.Next = mapIdxToNode[right]
			right--
		}
		cur = cur.Next
		turn = 1-turn
	}
	if cur != nil{
		cur.Next = nil
	}
}
