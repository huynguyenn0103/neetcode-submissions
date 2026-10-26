// O(l*l)
func handleWord(targets map[string]bool, str string)[]string{
	bytes := []byte(str)
	neighbors := make([]string, 0)
	for i := 0; i < len(bytes); i++ {
		for c := byte('a'); c <= 'z'; c++ {
			if c == str[i]{continue}
			bytes[i] = c
			newStr := string(bytes)
			if targets[newStr]{
				neighbors = append(neighbors, newStr)
			}
		}
		bytes[i]=str[i]
	}
	return neighbors
}

type QItem struct {
	val string
	level int
}

// Time Complexity:  l*l*n + n^2

func ladderLength(beginWord string, endWord string, wordList []string) int {
    words := make(map[string]bool)
	words[beginWord]=true
	for _, word := range wordList{
		words[word]=true
	}
	graph := make(map[string][]string, 0)
	graph[beginWord]=handleWord(words, beginWord)
	for _, word := range wordList{
		graph[word] = handleWord(words, word)
	}
	queue := make([]QItem, 0)
	queue = append(queue, QItem{val: beginWord, level: 0})
	visited := make(map[string]bool)
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		val, level := item.val, item.level
		if visited[val]{continue}
		if val == endWord{return level+1}
		visited[val]=true
		for _, neighbor := range graph[val]{
			queue = append(queue, QItem{val: neighbor, level: level+1})
		}
	}
	return 0
}
