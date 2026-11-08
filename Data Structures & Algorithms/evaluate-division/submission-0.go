type Node struct {
	node string
	weight float64
}

func calculate(query []string, graph map[string][]Node)float64{
	start, target := query[0],query[1]
	visited := make(map[string]bool)
	if graph[start] == nil {return -1.0}
	var dfs func(node string, multiply float64)float64
	dfs = func(node string, multiply float64)float64{
		visited[node]=true
		if node == target {
			return multiply
		}
		for _, neighbor := range graph[node]{
			neighborNode, weight := neighbor.node, neighbor.weight
			if visited[neighborNode]{continue}
			result := dfs(neighborNode, multiply * weight)
			if result != -1.0 {
				return result
			}
		}
		return -1.0
	}
	return dfs(start,1.0)
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    graph := make(map[string][]Node)
	for i, equation := range equations{
		from, to, weight := equation[0], equation[1], values[i]
		graph[from] = append(graph[from],Node{node: to, weight: weight})
		graph[to] = append(graph[to],Node{node:from, weight : 1.0/weight})
	}
	result := make([]float64,0)
	for _, query := range queries {
		result = append(result, calculate(query,graph))
	}
	return result
}
