/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func preorderTraversal(root *TreeNode) []int {
    result := make([]int,0)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode){
		if root == nil {return}
		result = append(result,root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return result
}

