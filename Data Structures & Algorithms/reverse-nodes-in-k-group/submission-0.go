/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	step := 0
	var ptrF, ptrL *ListNode
	ptrL = head
	for ptrL != nil {
		if step == k - 1 {
			ptrL = ptrL.Next
			var prev, cur, tmpF *ListNode
			if ptrF == nil {
				cur = head
			} else {
				cur = ptrF.Next
			}

			for cur != ptrL {
				tmp := cur.Next
				if prev == nil {
					cur.Next = ptrL
					tmpF = cur
				} else {
					cur.Next = prev
				}
				prev = cur
				cur = tmp
			}
			if ptrF != nil {
				ptrF.Next = prev
			} else {
				head = prev
			}
			if tmpF != nil {
				ptrF = tmpF
			}
			step = 0
		} 
		if ptrL != nil {
			ptrL = ptrL.Next
		}
		step++
	}
	return head
}

//1 -> 2 -> 3 -> 4 -> 5 -> 6
