/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		if p.Val != q.Val {return false}
		isSameLeft := isSameTree(p.Left, q.Left)
		isSameRight := isSameTree(p.Right, q.Right)
		if isSameLeft && isSameRight {
			return true
		}
		return false
	} else {
		return false
	}
}
