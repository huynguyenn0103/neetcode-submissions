/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func maxDepth(root *TreeNode) int {
    if root == nil {return 0}
	heightLeft := maxDepth(root.Left)
	heightRight := maxDepth(root.Right)
	return max(heightLeft, heightRight) + 1
}
