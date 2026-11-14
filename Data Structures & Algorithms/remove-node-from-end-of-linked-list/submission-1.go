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
	posFromStart := numOfNodes - n
	if posFromStart == 0 {
		head = head.Next
	} else {
		curPos := 0
		var prev *ListNode
		cur = head
		for curPos != posFromStart{
			prev = cur
			cur = cur.Next
			curPos++
		}
		prev.Next = cur.Next
	}
	return head
}
