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

func levelOrder(root *TreeNode) [][]int {
    result := make([][]int, 0)
	if root == nil {
		return result
	}
	currentLevel := -1
	queue := make([]Item, 0)
	queue = append(queue, Item{node: root, level: 0})
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		node, level := item.node, item.level
		if level != currentLevel {
			result = append(result, []int{})
			currentLevel++
		}
		result[currentLevel] = append(result[currentLevel], node.Val)
		if node.Left != nil {
			queue = append(queue, Item{node: node.Left, level: level + 1})
		}
		if node.Right != nil {
			queue = append(queue, Item{node: node.Right, level: level + 1})
		}
	}
	return result
}
