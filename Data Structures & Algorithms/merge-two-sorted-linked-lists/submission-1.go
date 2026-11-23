/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    var head, cur *ListNode
	cur1, cur2 := list1, list2
	for cur1 != nil && cur2 != nil {
		val := min(cur1.Val, cur2.Val)
		if cur1.Val >= cur2.Val{
			cur2 = cur2.Next
		} else {
			cur1 = cur1.Next
		}
		newNode := &ListNode{Val: val}
		if head == nil {
			head = newNode
			cur = newNode
		} else {
			cur.Next = newNode
			cur = cur.Next
		}
	}
	for cur1 != nil {
		newNode := &ListNode{Val: cur1.Val}
		if head == nil {
			head = newNode
			cur = newNode
		} else {
			cur.Next = newNode
			cur = cur.Next
		} 
		cur1 = cur1.Next
	}
	for cur2 != nil {
		newNode := &ListNode{Val: cur2.Val}
		if head == nil {
			head = newNode
			cur = newNode
		} else {
			cur.Next = newNode
			cur = cur.Next
		} 
		cur2 = cur2.Next
	}
	return head
}
