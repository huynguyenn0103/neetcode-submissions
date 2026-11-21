/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
    maxVal := -math.MaxInt
	var maxSumWithPathThroughNodeOrItsChild func(node *TreeNode)int
	maxSumWithPathThroughNodeOrItsChild =  func(node *TreeNode)int{
		if node == nil {return -math.MaxInt}
		maxSumRight := maxSumWithPathThroughNodeOrItsChild(node.Right)
		maxSumLeft := maxSumWithPathThroughNodeOrItsChild(node.Left)
		result := node.Val
		if maxSumRight > 0 {
			result = max(result, node.Val + maxSumRight)
		}
		if maxSumLeft > 0 {
			result = max(result, node.Val + maxSumLeft)
		}
		if maxSumRight > 0 && maxSumLeft > 0 {
			maxVal = max(maxVal, node.Val + maxSumRight + maxSumLeft)
		}
		maxVal = max(result, maxVal)
		return result
	}
	maxSumWithPathThroughNodeOrItsChild(root)
	return maxVal
}
