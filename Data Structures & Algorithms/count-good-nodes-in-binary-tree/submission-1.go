/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    var numOfGoodNodeWithCurrentMax func(node *TreeNode, curMax int)int
	numOfGoodNodeWithCurrentMax = func (node *TreeNode, curMax int)int{
		if node == nil {
			return 0
		}
		result := 0
		if node.Val >= curMax{
			result++
			curMax = node.Val
		}
		result += numOfGoodNodeWithCurrentMax(node.Left, curMax)
		result += numOfGoodNodeWithCurrentMax(node.Right, curMax)
		return result
	}
	return numOfGoodNodeWithCurrentMax(root, root.Val)
}
