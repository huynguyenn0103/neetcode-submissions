/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    var validateBST func(node *TreeNode, allowMin, allowMax int)bool
	validateBST = func(node *TreeNode, allowMin, allowMax int)bool{
		if node == nil {
			return true
		}
		if node.Left != nil && node.Val <= node.Left.Val {
			return false
		}
		if node.Right != nil && node.Val >= node.Right.Val{
			return false
		}
		if node.Val <= allowMin {
			return false
		}
		if node.Val >= allowMax{
			return false
		}
		return validateBST(node.Left, allowMin, node.Val) && validateBST(node.Right, node.Val, allowMax)
	}
	return validateBST(root, -math.MaxInt, math.MaxInt)
}
