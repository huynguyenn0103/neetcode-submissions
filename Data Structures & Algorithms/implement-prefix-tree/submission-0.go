type TreeNode struct {
	isEnd bool
	children map[byte]*TreeNode
}

type PrefixTree struct {
	root *TreeNode
}

func Constructor() PrefixTree {
    return PrefixTree{root: &TreeNode{
		children: make(map[byte]*TreeNode),
	}}
}

func (this *PrefixTree) Insert(word string) {
	node := this.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if node.children[c] == nil {
			node.children[c]=&TreeNode{
				children: make(map[byte]*TreeNode),
			}
		}
		node = node.children[c]
	}
	node.isEnd = true
}

func (this *PrefixTree) Search(word string) bool {
	node := this.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if node.children[c] == nil {return false}
		node = node.children[c]
	}
	return node.isEnd
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	node := this.root
	for i := 0; i < len(prefix); i++ {
		c := prefix[i]
		if node.children[c] == nil {return false}
		node = node.children[c]
	}
	return true
}	
