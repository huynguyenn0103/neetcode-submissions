const Visited = 2

func validTree(n int, edges [][]int) bool {
    graph := make([][]int, n)
	visited := make([]int, n)
	for node := range graph{
		graph[node]=make([]int, 0)
	}
	for _, edge := range edges{
		from, to := edge[0], edge[1]
		graph[from]=append(graph[from],to)
		graph[to]=append(graph[to],from)
	}
	var dfs func(node int, parent int)bool
	dfs = func(node int, parent int)bool {
		if visited[node]==Visited{
			return false
		}
		visited[node]=Visited
		for _, neighbor := range graph[node]{
			if neighbor == parent{continue}
			isSuccess := dfs(neighbor, node)
			if !isSuccess {
				return false
			}
		}
		return true
	}
	isSuccess := dfs(0,-1)
	if !isSuccess{
		return false
	}
	for _, status := range visited {
		if status == 0 {
			return false
		}
	}
	return true
}
