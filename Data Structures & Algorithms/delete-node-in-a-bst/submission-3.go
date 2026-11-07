func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return nil
    }

    // 1. Search for the node
    if key < root.Val {
        root.Left = deleteNode(root.Left, key)
    } else if key > root.Val {
        root.Right = deleteNode(root.Right, key)
    } else {
        // 2. Node found

        // Case 1: No children
        if root.Left == nil && root.Right == nil {
            return nil
        }

        // Case 2: Only right child
        if root.Left == nil {
            return root.Right
        }

        // Case 2: Only left child
        if root.Right == nil {
            return root.Left
        }

        // Case 3: Two children
        successor := root.Right

        for successor.Left != nil {
            successor = successor.Left
        }

        root.Val = successor.Val

        root.Right = deleteNode(root.Right, successor.Val)
    }

    return root
}