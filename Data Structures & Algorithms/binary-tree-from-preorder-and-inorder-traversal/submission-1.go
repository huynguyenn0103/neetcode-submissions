/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// PreOrder: NLR => 1,2,3,4
// InOrder: LNR => 2,1,3,4.,m

func buildTree(preorder []int, inorder []int) *TreeNode {
    mapPreorderValToIdx := make(map[int]int)
	for i, num := range preorder{
		mapPreorderValToIdx[num]=i
	}
	var  buildNodeFromInOrder func(inorder []int) *TreeNode
	buildNodeFromInOrder = func(inorder []int) *TreeNode{
		if len(inorder) == 0 {
			return nil
		}
		breakIdx := 0
		minIdx := math.MaxInt
		for i, num := range inorder{
			if mapPreorderValToIdx[num] < minIdx{
				minIdx = mapPreorderValToIdx[num]
				breakIdx = i
			}
		}
		node := &TreeNode{Val: inorder[breakIdx]}
		node.Left = buildNodeFromInOrder(inorder[0:breakIdx])
		node.Right = buildNodeFromInOrder(inorder[breakIdx+1:])
		return node
	}
	return buildNodeFromInOrder(inorder)
}
