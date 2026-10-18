/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
  if head == nil || head.Next == nil {return false}
  slowPtr, fastPtr := head, head.Next
  for fastPtr != nil && fastPtr.Next != nil {
	if slowPtr == fastPtr{return true}
	slowPtr = slowPtr.Next
	fastPtr = fastPtr.Next.Next
  }  
  return false
}
