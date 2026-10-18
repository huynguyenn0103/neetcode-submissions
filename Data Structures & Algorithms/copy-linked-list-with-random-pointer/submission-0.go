/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

const IndexOfNull = -1

func copyRandomList(head *Node) *Node {
    mapPointerToIndex := make(map[*Node]int)
	curPtr, curIdx := head, 0
	for curPtr != nil {
		mapPointerToIndex[curPtr] = curIdx
		curPtr = curPtr.Next
		curIdx++
	}
	randomIndexList := make([]int, 0)
	curPtr = head
	for curPtr != nil {
		if curPtr.Random == nil {
			randomIndexList = append(randomIndexList, IndexOfNull)
		} else {
			randomIndexList = append(randomIndexList, mapPointerToIndex[curPtr.Random])
		}
		curPtr = curPtr.Next
	}
	newNodesList := make([]*Node, 0)
	var rHead, rCurPtr *Node
	curPtr = head
	for curPtr != nil {
		newNode := &Node{Val: curPtr.Val, Next: nil}
		if rHead == nil {
			rHead = newNode
			rCurPtr = newNode
		} else {
			rCurPtr.Next = newNode
			rCurPtr = rCurPtr.Next
		}
		curPtr = curPtr.Next
		newNodesList = append(newNodesList, rCurPtr)
	}
	rCurPtr = rHead
	curIdx = 0
	for rCurPtr != nil {
		randomIdx := randomIndexList[curIdx]
		if randomIdx != IndexOfNull{
			rCurPtr.Random = newNodesList[randomIdx]
		}
		curIdx++
		rCurPtr = rCurPtr.Next
	}
	return rHead
}
