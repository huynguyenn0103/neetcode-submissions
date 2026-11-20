/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	var containNode func(node, target *TreeNode)bool
	containNode = func(node, target *TreeNode)bool{
		if node == nil {
			return false
		}
		if node.Val == target.Val{
			return true
		}
		if node.Val > target.Val{
			return containNode(node.Left, target)
		}
		return containNode(node.Right, target)
	}

	var check func(node *TreeNode) *TreeNode
	check = func(node *TreeNode) *TreeNode{
		if node == nil{
			return nil
		}
		if containNode(node.Left, p) && containNode(node.Left, q){
			return check(node.Left)
		}
		if containNode(node.Right, p) && containNode(node.Right, q){
			return check(node.Right)
		}
		if containNode(node,p) && containNode(node,q){
			return node
		}
		return nil
	}
	return check(root)
}
