func isValidBST(root *TreeNode) bool {
    var validate func(node *TreeNode, minVal, maxVal int64) bool

    validate = func(node *TreeNode, minVal, maxVal int64) bool {
        if node == nil {
            return true
        }

        val := int64(node.Val)
        if val <= minVal || val >= maxVal {
            return false
        }

        return validate(node.Left, minVal, val) &&
            validate(node.Right, val, maxVal)
    }

    return validate(root, math.MinInt64, math.MaxInt64)
}