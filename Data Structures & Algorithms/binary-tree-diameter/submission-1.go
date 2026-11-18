/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func heightOfNode(root *TreeNode) int{
	if root == nil {
		return 0
	}
	return 1 + max(heightOfNode(root.Left),heightOfNode(root.Right))
}

func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
		return 0
	}
	result := heightOfNode(root.Left) + heightOfNode(root.Right)
	result = max(result, diameterOfBinaryTree(root.Left))
	result = max(result, diameterOfBinaryTree(root.Right))
	return result
}
