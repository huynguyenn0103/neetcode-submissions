type TrieNode struct{
	children [26]*TrieNode
	isWord bool
}

type Trie struct{
	root *TrieNode
}

func NewTrie() Trie{
	return Trie{root: &TrieNode{}}
}

func (t *Trie) AddWord(word string){
	n := len(word)
	cur := t.root
	for i := 0; i < n; i++ {
		c := word[i]
		idx := c-'a'
		if cur.children[idx] == nil{
			cur.children[idx]=&TrieNode{}
		} 
		cur = cur.children[idx]
	}
	cur.isWord = true
}

func (t *Trie) Search(word string)bool{
	n := len(word)
	var isValid func(idx int, node *TrieNode)bool
	isValid = func(idx int, node *TrieNode)bool{
		if idx == n {
			if node == nil {
				return false
			}
			return node.isWord
		}
		if node == nil {
			return false
		}
		c := word[idx]
		if word[idx] != '.'{
			if node.children[c-'a'] == nil {
				return false
			}
			return isValid(idx+1, node.children[c-'a'])
		} else {
			result := false
			for i := byte('a'); i < 'z'; i++ {
				result = result || isValid(idx+1, node.children[i-'a'])
			}
			return result
		}
	}
	return isValid(0, t.root)
}


type WordDictionary struct {
    trie Trie
}

func Constructor() WordDictionary {
    return WordDictionary{
		trie: NewTrie(),
	}
}

func (this *WordDictionary) AddWord(word string)  {
    this.trie.AddWord(word)
}

func (this *WordDictionary) Search(word string) bool {
	return this.trie.Search(word)
}
