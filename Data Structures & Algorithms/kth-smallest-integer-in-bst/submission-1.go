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
	count := 0
	var traverseInorder func(root *TreeNode)
	traverseInorder = func(root *TreeNode){
		if root == nil{return}
		traverseInorder(root.Left)
		count++
		if count == k{
			result = root.Val
		}
		traverseInorder(root.Right)
	}
	traverseInorder(root)
	return result
}
