/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    lists := make([]*ListNode, 0)
	cur := head
	for cur != nil {
		lists = append(lists, cur)
		cur = cur.Next
	}
	left, right := 0, len(lists)-1
	turn := 0

	cur = nil
	for left <= right {
		if turn == 0 {
			if cur == nil {
				cur = lists[left]
			} else {
				cur.Next = lists[left]
				cur = cur.Next
			}
			left++
		} else {
			cur.Next = lists[right]
			cur = cur.Next
			right--
		}
		turn = 1 - turn
	}
	cur.Next = nil
}
