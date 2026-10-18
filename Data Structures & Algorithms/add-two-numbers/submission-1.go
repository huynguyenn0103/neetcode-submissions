/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverse(l *ListNode) (int){
	count := 0
	var cur *ListNode
	cur = l
	for cur != nil {
		count++
		cur = cur.Next
	}
	return count
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    len1 := reverse(l1)
	len2 := reverse(l2)
	// Make sure l1 >= l2
	if len2 > len1 {
		l1, l2 = l2, l1
	}
	var prevCur1 *ListNode
	cur1, cur2 := l1, l2
	carry := 0
	for cur1 != nil && cur2 != nil {
		total := cur1.Val + cur2.Val + carry
		carry = total / 10
		total = total % 10
		cur1.Val = total
		prevCur1 = cur1
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	for carry == 1 {
		if cur1 == nil {
			prevCur1.Next = &ListNode{Val: 1, Next: nil}
			break
		} else {
			newVal := cur1.Val + 1
			if newVal == 10 {
				newVal = 0
				carry = 1
			} else {
				carry = 0
			}
			prevCur1 = cur1
			cur1.Val = newVal
			cur1 = cur1.Next
		}
	}
	return l1
}
