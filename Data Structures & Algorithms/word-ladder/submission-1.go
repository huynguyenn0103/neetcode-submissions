type Item struct{
	val string
	level int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
    mapWord := make(map[string]bool)
	for _, word := range wordList{
		mapWord[word]=true
	}
	visited := make(map[string]bool)
	queue := make([]Item,0)
	queue = append(queue, Item{val: beginWord, level: 1})
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		val, level := item.val, item.level
		if visited[val]{continue}
		if val == endWord{
			return level
		}
		visited[val]=true
		n := len(val)
		for i := 0; i < n; i++ {
			tmp := []byte(val)
			for c := byte('a'); c <= 'z'; c++{
				tmp[i]=c
				newStr := string(tmp)
				if mapWord[newStr] && !visited[newStr]{
					queue = append(queue, Item{val:newStr,level:level+1})
				}
			}
		}
	}
	return 0
}
