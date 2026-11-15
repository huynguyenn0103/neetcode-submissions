/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry := 0
	cur1, cur2 := l1, l2
	var result, curR *ListNode
	for cur1 != nil && cur2 != nil{
		sum := (cur1.Val + cur2.Val + carry) % 10
		carry = (cur1.Val + cur2.Val + carry) / 10
		newNode := &ListNode{Val: sum}
		if result == nil {
			result = newNode
			curR = result
		} else {
			curR.Next = newNode
			curR = curR.Next 
		}
		cur1 = cur1.Next
		cur2 = cur2.Next
	}

	for cur1 != nil {
		sum := (cur1.Val + carry) % 10
		carry = (cur1.Val + carry) / 10
		curR.Next = &ListNode{Val: sum}
		curR = curR.Next
		cur1 = cur1.Next
	}

	for cur2 != nil {
		sum := (cur2.Val + carry) % 10
		carry = (cur2.Val + carry) / 10
		curR.Next = &ListNode{Val: sum}
		curR = curR.Next
		cur2 = cur2.Next
	}

	if carry == 1{
		curR.Next = &ListNode{Val: 1}
	}
	return result
}
