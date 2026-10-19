/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func height(root *TreeNode) int{
	if root == nil {return 0}
	heightLeft := height(root.Left)
	heightRight := height(root.Right)
	return max(heightLeft, heightRight) + 1
}


func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	heightLeft := height(root.Left)
	heightRight := height(root.Right)
	if heightLeft - heightRight > 1 || heightRight - heightLeft > 1{
		return false
	} else {
		isRightBalanced := isBalanced(root.Left)
		isLeftBalanced := isBalanced(root.Right)
		if !isRightBalanced || !isLeftBalanced {
			return false
		} 
		return true
	}
}
