func buildEdge(w1, w2 string) ([]byte,bool){
	n1, n2 := len(w1), len(w2)
	ptr1, ptr2 := 0, 0
	for ptr1 < n1 && ptr2 < n2{
		if w1[ptr1] != w2[ptr2]{
			return []byte{w1[ptr1],w2[ptr2]}, true
		}
		ptr1++
		ptr2++
	}
	if n1 <= n2 {
		return []byte{}, true
	}
	return []byte{}, false
} 

func foreignDictionary(words []string) string {
    graph := make(map[byte][]byte)
	indegrees := make(map[byte]int)
	n := len(words)
	for i := 1; i < n; i++{
		edge, ok := buildEdge(words[i-1],words[i])
		if !ok {return ""}
		if len(edge) == 0 {continue}
		from, to := edge[0], edge[1]
		if graph[from] == nil {
			graph[from] = make([]byte,0)
		}
		if graph[to] == nil {
			graph[to] = make([]byte,0)
		}
		graph[from] = append(graph[from],to)
		indegrees[to]++
	}
	for _, word := range words{
		l := len(word)
		for i := 0; i < l; i++{
			c := word[i]
			if graph[c] == nil {
				graph[c] = make([]byte,0)
			}
		}
	}
	result := make([]byte,0)
	queue := make([]byte, 0)
	for node := range graph{
		if indegrees[node] == 0{
			queue = append(queue, node)
		}
	}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)
		for _, neighborNode := range graph[node]{
			indegrees[neighborNode]--
			if indegrees[neighborNode]==0{
				queue = append(queue, neighborNode)
			}
		}
	}

	if len(result) == len(graph){
		return string(result)
	}
	return ""
}
