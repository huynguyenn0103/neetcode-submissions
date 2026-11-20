/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func abs(x int)int{
	if x >= 0 {return x}
	return -x
}

func isBalanced(root *TreeNode) bool {
    isBalance := true
	var calHeight func(node *TreeNode)int
	calHeight = func(node *TreeNode) int{
		if node == nil {
			return 0
		}
		heightLeft := calHeight(node.Left)
		heightRight := calHeight(node.Right)
		if isBalance {
			if abs(heightLeft - heightRight) > 1{
				isBalance = false
			}
		}
		return 1 + max(heightLeft, heightRight)
	}
	calHeight(root)
	return isBalance
}
