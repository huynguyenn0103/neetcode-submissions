/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    var isSameTree func(root *TreeNode, subRoot *TreeNode)bool
	isSameTree = func(root *TreeNode, subRoot *TreeNode)bool{
		if root == nil && subRoot == nil {
			return true
		}
		if root == nil || subRoot == nil {
			return false
		}
		if root.Val != subRoot.Val {
			return false
		}
		return isSameTree(root.Left, subRoot.Left) && isSameTree(root.Right, subRoot.Right)
	}

	var subTree func(root *TreeNode, subRoot *TreeNode)bool
	subTree = func(root *TreeNode, subRoot *TreeNode)bool{
		if isSameTree(root, subRoot){
			return true
		}
		if root == nil {return false}
		return subTree(root.Left, subRoot) || subTree(root.Right, subRoot)
	}
	return subTree(root,subRoot)
}
