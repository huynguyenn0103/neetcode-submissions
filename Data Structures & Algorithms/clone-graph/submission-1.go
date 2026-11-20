/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil{
		return nil
	}
    mapValToNode := make(map[int]*Node)
	firstNode := &Node{Val: node.Val}
	var traverse func(orgNode *Node, newNode *Node)
	traverse = func(orgNode *Node,newNode *Node){
        mapValToNode[newNode.Val]=newNode
		neighbors := make([]*Node,0)
		for _, neighborNode := range orgNode.Neighbors{
            val := neighborNode.Val
			if mapValToNode[val] != nil{
               neighbors = append(neighbors, mapValToNode[val])
			} else{
				cloneNode := &Node{Val: val}
				traverse(neighborNode, cloneNode)
				neighbors = append(neighbors, cloneNode)
			}
		}
		newNode.Neighbors = neighbors
	}
	traverse(node, firstNode)
	return firstNode
}
