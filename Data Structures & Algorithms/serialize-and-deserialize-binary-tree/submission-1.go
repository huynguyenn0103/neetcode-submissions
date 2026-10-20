type Codec struct{}

func Constructor() Codec {
    return Codec{}
}

func (this *Codec) serialize(root *TreeNode) string {
    var builder strings.Builder

    var dfs func(*TreeNode)

    dfs = func(node *TreeNode) {
        if node == nil {
            builder.WriteString("#,")
            return
        }

        builder.WriteString(strconv.Itoa(node.Val))
        builder.WriteString(",")

        dfs(node.Left)
        dfs(node.Right)
    }

    dfs(root)

    return builder.String()
}

func (this *Codec) deserialize(data string) *TreeNode {
    values := strings.Split(data, ",")
    idx := 0

    var dfs func() *TreeNode

    dfs = func() *TreeNode {
        if values[idx] == "#" {
            idx++
            return nil
        }

        val, _ := strconv.Atoi(values[idx])
        idx++

        root := &TreeNode{Val: val}

        root.Left = dfs()
        root.Right = dfs()

        return root
    }

    return dfs()
}