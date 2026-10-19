/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func check(root *TreeNode, subRoot *TreeNode) bool{
	if root == nil && subRoot == nil {
		return true
	} else if root != nil && subRoot != nil {
		if root.Val != subRoot.Val {
			return false
		}		
		isLeftSame := check(root.Left, subRoot.Left)
		isRightSame := check(root.Right, subRoot.Right)
		if isLeftSame && isRightSame {
			return true
		}
		return false
	} else {
		return false
	}
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if check(root, subRoot){
		return true
	} else {
		if root == nil {return false}
		checkLeft := isSubtree(root.Left, subRoot)
		checkRight := isSubtree(root.Right, subRoot)
		if checkLeft || checkRight{return true}
		return false
	}
}
