/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseIfEnoughK(beforeHead *ListNode, head *ListNode, k int) *ListNode{
	count := 0
	cur := head
	for cur != nil && count < k{
		cur = cur.Next
		count++
	}
	if count < k {
		return head
	}
	remainHead := cur
	cur = head
	prev := remainHead
	for cur != nil && k > 0 {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
		k--
	}
	if beforeHead != nil {
		beforeHead.Next = prev
	}
	return prev
}

func reverseKGroup(head *ListNode, k int) *ListNode {
    var newHead, prev *ListNode
	cur := head
	for cur != nil {
		count := 0
		node := reverseIfEnoughK(prev, cur, k)
		if newHead == nil{
			newHead = node
		}
		cur = node
		for cur != nil && count < k{
			prev = cur
			cur = cur.Next
			count++
		}
	}
	return newHead
}
