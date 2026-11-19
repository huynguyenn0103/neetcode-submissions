/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    mapNodeToIdx := make(map[*Node]int)
	cur := head
	curIdx := 0
	for cur != nil {
		mapNodeToIdx[cur]=curIdx
		curIdx++
		cur = cur.Next
	}

	mapIdxToNode := make(map[int]*Node)
	var newHead, newCur *Node
	for i := 0; i < curIdx; i++ {
		newNode := &Node{}
		if newHead == nil {
			newHead = newNode
			newCur = newHead
		} else {
			newCur.Next = newNode
			newCur = newCur.Next
		}
		mapIdxToNode[i]=newNode
	}

	newCur = newHead
	cur = head
	for newCur != nil{
		newCur.Val = cur.Val
		if cur.Random != nil {
			idxOfRandom := mapNodeToIdx[cur.Random]
			newNodeInIdx := mapIdxToNode[idxOfRandom]
			newCur.Random = newNodeInIdx
		}
		newCur = newCur.Next
		cur = cur.Next
	}
	return newHead
}
