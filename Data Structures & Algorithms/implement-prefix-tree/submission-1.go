type Node struct{
	isEnd bool
	children map[byte]*Node
}

type PrefixTree struct {
	root *Node
}

func Constructor() PrefixTree {
    return PrefixTree{root: &Node{children: make(map[byte]*Node)}}
}

func (this *PrefixTree) Insert(word string) {
	cur := this.root
	n := len(word)
	for i := 0; i < n; i++{
		c := word[i]
		if cur.children[c] == nil {
			cur.children[c] = &Node{children: make(map[byte]*Node)}
		}
		cur = cur.children[c]
	}
	cur.isEnd = true
}

func (this *PrefixTree) Search(word string) bool {
	cur := this.root
	n := len(word)
	for i := 0; i < n; i++{
		c := word[i]
		if cur.children[c] == nil {return false}
		cur = cur.children[c]
	}
	return cur.isEnd
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	cur := this.root
	n := len(prefix)
	for i := 0; i < n; i++{
		c := prefix[i]
		if cur.children[c] == nil {return false}
		cur = cur.children[c]
	}
	return true
}
