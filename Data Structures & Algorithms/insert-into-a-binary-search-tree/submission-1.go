/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
    cur := root
	prev := cur
	for cur != nil {
		prev = cur
		if val >= cur.Val {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}
	if val >= prev.Val {
		prev.Right = &TreeNode{Val: val}
	} else {
		prev.Left = &TreeNode{Val: val}
	}
	return root
}
