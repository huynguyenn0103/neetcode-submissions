type TreeNode struct {
	isEnd bool
	children map[byte]*TreeNode
}

type WordDictionary struct {
    root *TreeNode
}

func Constructor() WordDictionary {
    return WordDictionary {
		root: &TreeNode{
			children: make(map[byte]*TreeNode),
		},
	}
}

func (this *WordDictionary) AddWord(word string)  {
    node := this.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if node.children[c] == nil {
			node.children[c]=&TreeNode{children: make(map[byte]*TreeNode)}
		}
		node = node.children[c]
	}
	node.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	var dfs func(node *TreeNode, word string)bool 
	dfs = func(node *TreeNode, word string)bool {
		if len(word) == 0 {
			return node.isEnd
		}
		c := word[0]
		if c != '.'{
			if node.children[c] == nil {return false}
			return dfs(node.children[c], word[1:len(word)])
		} else {
			for _, childNode := range node.children{
				isSuccess := dfs(childNode, word[1:len(word)])
				if isSuccess{
					return true
				}
			}
		}
		return false
	}

	return dfs(this.root, word)
}
