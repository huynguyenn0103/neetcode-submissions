/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func height(root *TreeNode)int {
	if root == nil {return 0}
	heightLeft := height(root.Left)
	heightRight := height(root.Right)
	return max(heightLeft, heightRight) + 1
}

func diameterThroughRoot(root *TreeNode) int {
	heightLeft := height(root.Left)
	heightRight := height(root.Right)
	d := 0
	if root.Left != nil {
		d += heightLeft
	}
	if root.Right != nil {
		d += heightRight
	}
	return d
}

func diameterOfBinaryTree(root *TreeNode) int {
    result := 0
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode){
		if root == nil {return}
		result = max(result, diameterThroughRoot(root))
		traverse(root.Left)
		traverse(root.Right)
	}
	traverse(root)
	return result
}
