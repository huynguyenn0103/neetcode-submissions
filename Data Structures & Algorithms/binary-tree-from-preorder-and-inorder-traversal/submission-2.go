func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder)==0 || len(inorder)==0{
		return nil
	}
    root:=&TreeNode{}
	root.Val=preorder[0]
	mid:=0
	for i,val:=range inorder{
		if val==root.Val{
			mid=i
			break
		}
	}
	root.Left=buildTree(preorder[1:mid+1],inorder[:mid])
	root.Right=buildTree(preorder[mid+1:],inorder[mid+1:])

	return root
	
}