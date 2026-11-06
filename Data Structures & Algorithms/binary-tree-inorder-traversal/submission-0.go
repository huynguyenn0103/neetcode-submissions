/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Root+Left+Right

func inorderTraversal(root *TreeNode) []int {
	result := make([]int,0)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode){
		if root == nil {return}
		dfs(root.Left)
		result = append(result,root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return result
}
