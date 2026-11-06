/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal(root *TreeNode) []int {
	result := make([]int,0)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode){
		if root == nil {return}
		dfs(root.Left)
		dfs(root.Right)
		result = append(result,root.Val)
	}
	dfs(root)
	return result
    
}
