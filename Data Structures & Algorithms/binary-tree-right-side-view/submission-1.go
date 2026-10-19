/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Item struct {
	node *TreeNode
	level int
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
    nodesByLevel := make([]Item, 0)
	queue := make([]Item, 0)
	queue = append(queue, Item{node: root, level: 0})
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		if len(nodesByLevel) > 0 && nodesByLevel[len(nodesByLevel)-1].level == item.level {
			nodesByLevel = nodesByLevel[:len(nodesByLevel)-1]
		}
		nodesByLevel = append(nodesByLevel, item)
		node, level := item.node, item.level
		if node.Left != nil {
			queue = append(queue, Item{node: node.Left, level: level + 1})
		}
		if node.Right != nil {
			queue = append(queue, Item{node: node.Right, level: level + 1})
		}
	}
	result := make([]int, 0)
	for _, nodeItem := range nodesByLevel{
		result = append(result, nodeItem.node.Val)
	}
	return result
}
