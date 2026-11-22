

type Codec struct {}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    result := make([]string, 0)

    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            result = append(result, "null")
            return
        }

        result = append(result, strconv.Itoa(node.Val))

        dfs(node.Left)
        dfs(node.Right)
    }

    dfs(root)

    return strings.Join(result, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
    values := strings.Split(data, ",")
    idx := 0

    var dfs func() *TreeNode
    dfs = func() *TreeNode {
        if values[idx] == "null" {
            idx++
            return nil
        }

        val, _ := strconv.Atoi(values[idx])
        idx++

        node := &TreeNode{Val: val}

        node.Left = dfs()
        node.Right = dfs()

        return node
    }

    return dfs()
}