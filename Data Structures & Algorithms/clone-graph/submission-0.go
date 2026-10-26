/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
   if node == nil {return nil}
   mapValToNewNode := make(map[int]*Node)
   visited := make(map[*Node]bool)
   queue := make([]*Node, 0)
   queue = append(queue, node)
   for len(queue) > 0 {
	  node := queue[0]
	  queue = queue[1:]
	  if visited[node]{continue}
	  visited[node]=true
	  newNode := &Node{Val: node.Val}
	  if mapValToNewNode[node.Val] != nil {
		newNode = mapValToNewNode[node.Val]
	  } else {
		mapValToNewNode[node.Val]=newNode
	  }
	  neighbors := make([]*Node, 0)
	  for _, neighborNode := range node.Neighbors{
		 queue = append(queue, neighborNode)
		 newNeighborNode := &Node{Val: neighborNode.Val}
		 if mapValToNewNode[neighborNode.Val] != nil {
			newNeighborNode = mapValToNewNode[neighborNode.Val]
		 } else {
			mapValToNewNode[neighborNode.Val]=newNeighborNode
		 }
		neighbors = append(neighbors, newNeighborNode)
	  }
	  newNode.Neighbors = neighbors
    }
	return mapValToNewNode[node.Val]
}
