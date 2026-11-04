/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func gcd(val1, val2 int)int{
	for val1 != val2 {
		if val1 > val2 {
			val1 = val1 - val2
		} else {
			val2 = val2 - val1
		}
	}
	return val1
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
    prev, cur := head, head.Next
	for cur != nil {
		newNode := &ListNode{Val: gcd(prev.Val,cur.Val)}
		prev.Next = newNode
		newNode.Next=cur
		prev = cur
		cur = cur.Next
	}
	return head
}
