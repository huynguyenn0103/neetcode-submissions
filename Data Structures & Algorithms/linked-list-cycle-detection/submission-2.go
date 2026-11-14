/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	if head == nil {return false}
    slowPtr, fastPtr := head, head.Next
	for slowPtr != nil && fastPtr != nil{
		if slowPtr == fastPtr {return true}
		slowPtr = slowPtr.Next
		if fastPtr.Next == nil{
			return false
		}
		fastPtr = fastPtr.Next.Next
	}
	return false
}
