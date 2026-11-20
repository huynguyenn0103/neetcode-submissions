func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    node := root

    for node != nil {
        if p.Val < node.Val && q.Val < node.Val {
            node = node.Left
        } else if p.Val > node.Val && q.Val > node.Val {
            node = node.Right
        } else {
            return node
        }
    }

    return nil
}