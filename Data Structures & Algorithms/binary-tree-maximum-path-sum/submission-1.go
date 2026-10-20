/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
const Infinity = math.MaxInt
func maxPathSum(root *TreeNode) int {
	result := -Infinity
	var buildMax func(root *TreeNode)int
	buildMax = func(root *TreeNode)int {
		if root == nil {return 0}
		leftVal := buildMax(root.Left)
		rightVal := buildMax(root.Right)
		newVal := max(leftVal + root.Val, rightVal + root.Val, root.Val)
		result = max(result, newVal, leftVal + rightVal + root.Val)
		return newVal
	}
	buildMax(root)
	return result
}
