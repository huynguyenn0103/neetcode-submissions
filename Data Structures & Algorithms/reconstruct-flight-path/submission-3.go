import "slices"

func findItinerary(tickets [][]string) []string {
    result := make([]string,0)
	graph := make(map[string][]string)
	for _, ticket := range tickets{
		from, to := ticket[0],ticket[1]
		graph[from]=append(graph[from],to)
	}
	for node := range graph {
		slices.Sort(graph[node])
		slices.Reverse(graph[node])
	}

	var dfs func(string)
	dfs = func(node string){
		for len(graph[node]) > 0 {
			n := len(graph[node])
			next := graph[node][n-1]
			graph[node]=graph[node][:n-1]
			dfs(next)
		}
		result = append(result, node)
	}
	dfs("JFK")
	slices.Reverse(result)
	return result
}
