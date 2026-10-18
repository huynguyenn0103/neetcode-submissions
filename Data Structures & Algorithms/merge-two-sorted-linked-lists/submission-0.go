/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var prevCur1 *ListNode
	cur1, cur2 := list1, list2
	result := list1
	for cur1 != nil && cur2 != nil {
		if cur1.Val < cur2.Val{
			prevCur1 = cur1
			cur1 = cur1.Next
		} else {
			tmp := cur2.Next
			cur2.Next = cur1
			if prevCur1 == nil{
				result = cur2
			} else {
				prevCur1.Next = cur2
			}
			prevCur1 = cur2
			cur2 = tmp
		}
	}
	if cur2 != nil {
		if prevCur1 == nil {
			return cur2
		} else {
			prevCur1.Next = cur2
		}
	}
	return result
}
