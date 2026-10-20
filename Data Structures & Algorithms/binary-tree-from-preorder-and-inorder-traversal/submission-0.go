/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
   positionNodeInOrder := make(map[int]int)
   for i := 0; i < len(inorder); i++ {
	positionNodeInOrder[inorder[i]]=i
   } 
   var head *TreeNode
   for i := 0; i < len(preorder); i++ {
		if head == nil {
			head = &TreeNode{Val: preorder[i]}
		} else {
			cur := head
			for {
				newVal := preorder[i]
				if positionNodeInOrder[newVal] < positionNodeInOrder[cur.Val]{
					if cur.Left == nil {
						cur.Left = &TreeNode{Val: newVal}
						break
					} else {	
						cur = cur.Left
					}
				} else {
					if cur.Right == nil {
						cur.Right = &TreeNode{Val: newVal}
						break
					} else {
						cur = cur.Right
					}
				}	
			}
		}
    }
	return head
}
