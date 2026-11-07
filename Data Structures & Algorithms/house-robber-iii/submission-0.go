/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func rob(root *TreeNode) int {
	cache := make(map[*TreeNode]int)

    var maxProfitFromSubtree func(r *TreeNode)int
	maxProfitFromSubtree = func(r *TreeNode)int{
		if r == nil {
			return 0
		}
		if val, existed := cache[r]; existed{
			return val
		}
		var maxResult int

		maxResult = r.Val
		if r.Left != nil{
			maxResult += maxProfitFromSubtree(r.Left.Left)
			maxResult += maxProfitFromSubtree(r.Left.Right)
		}
		if r.Right != nil {
			maxResult += maxProfitFromSubtree(r.Right.Left)
			maxResult += maxProfitFromSubtree(r.Right.Right)
		}
		maxResult = max(maxResult, maxProfitFromSubtree(r.Left) + maxProfitFromSubtree(r.Right))
		cache[r]=maxResult
		return cache[r]
	}
	return maxProfitFromSubtree(root)
}
