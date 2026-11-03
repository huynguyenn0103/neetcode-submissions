/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	left, right = left-1, right-1
	var startPtr, endPtr, leftPtr, rightPtr ,prev *ListNode
	cur := head
	count := 0
	for cur != nil {
		if count == left-1{
			startPtr = cur
		}
		if count == right+1{
			endPtr = cur
		}
		if count == left {
			leftPtr = cur
		}
		if count == right {
			rightPtr = cur
		}
		if count > left && count <= right {
			tmp := cur.Next
			cur.Next = prev
			prev = cur
			cur = tmp
		} else {
			prev = cur
			cur = cur.Next
		}
		count++
	}
	if startPtr != nil {
		startPtr.Next = rightPtr
	} else {
		head = rightPtr
	}
	if leftPtr != nil {
		leftPtr.Next = endPtr
	}
	return head
}
