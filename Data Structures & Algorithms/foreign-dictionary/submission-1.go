

func foreignDictionary(words []string)string {
    graph := make(map[byte][]byte)
	indegrees := make(map[byte]int)
	if len(words) == 1{
		return words[0]
	} 
	compare := func(s1, s2 string) bool{
		ptr := 0
		hasEdge := false
		for ptr < len(s1) && ptr < len(s2){
			if graph[s1[ptr]] == nil {
				graph[s1[ptr]]=make([]byte,0)
			}
			if graph[s2[ptr]] == nil {
				graph[s2[ptr]]=make([]byte,0)
			}
			if s1[ptr] != s2[ptr] && !hasEdge{
				hasEdge = true
				graph[s1[ptr]]=append(graph[s1[ptr]],s2[ptr])
				indegrees[s2[ptr]]++
			}
			ptr++
		}
		if !hasEdge && len(s1) > len(s2) {return false}
		for ptr < len(s1){
			if graph[s1[ptr]] == nil {
				graph[s1[ptr]]=make([]byte,0)
			}
			ptr++
		}
		for ptr < len(s2){
			if graph[s2[ptr]] == nil{
				graph[s2[ptr]]=make([]byte,0)
			}
			ptr++
		}
		return true
	}
	for i := 1; i < len(words); i++ {
		isOk := compare(words[i-1],words[i])
		if !isOk{return ""}
	}
	queue := make([]byte, 0)
	for node := range graph{
		if indegrees[node]==0{
			queue = append(queue, node)
		}
	}
	result := make([]byte,0)
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		result = append(result, c)
		for _, neighbor := range graph[c]{
			indegrees[neighbor]--
			if indegrees[neighbor]==0{
				queue = append(queue, neighbor)
			}
		}
	}
	if len(result) != len(graph){
		return ""
	}
	return string(result)
}
