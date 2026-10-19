/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	var result int
    var traverse func(root *TreeNode)
	traverse = func(root *TreeNode){
		if root == nil {
			return
		}
		traverse(root.Left)
		if k == 0 {return}
		result = root.Val
		k--
		if k == 0 {return}
		traverse(root.Right)
	}
	traverse(root)
	return result
}
