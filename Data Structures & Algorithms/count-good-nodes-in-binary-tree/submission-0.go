/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
const Infinity = math.MaxInt
func goodNodes(root *TreeNode) int {
    result := 0
	var dfs func(root *TreeNode, prevMax int)
	dfs = func(root *TreeNode, prevMax int){
		if root == nil {
			return
		}
		if root.Val >= prevMax {
			result++
			prevMax = root.Val
		}
		dfs(root.Left, prevMax)
		dfs(root.Right, prevMax)
	}
	dfs(root, -Infinity)
	return result
}
