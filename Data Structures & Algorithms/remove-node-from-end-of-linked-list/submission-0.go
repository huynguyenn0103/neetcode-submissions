/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    numOfNodes := 0
	cur := head
	for cur != nil {
		numOfNodes++
		cur = cur.Next
	}
	idxToRemove := numOfNodes - n
	var prev *ListNode
	cur = head
	idx := 0
	for idx != idxToRemove {
		prev = cur
		cur = cur.Next
		idx++
	}
	if prev == nil {
		head = head.Next
	} else {
		prev.Next = cur.Next
		cur.Next = nil
	}
	return head
}
