func buildTopologicalSort(edges [][]int, n int)(map[int]int,bool){
	graph := make([][]int, n)
	indegrees := make([]int, n)
	for node := range graph{
		graph[node]=make([]int,0)
	}
	for _, edge := range edges{
		from, to := edge[0]-1, edge[1]-1
		graph[from]=append(graph[from],to)
		indegrees[to]++
	}
	queue := make([]int,0)
	for node, degree := range indegrees{
		if degree == 0 {
			queue = append(queue, node)
		}
	}

	result := make([]int,0)
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		result = append(result, item + 1)
		for _, neighbor := range graph[item]{
			indegrees[neighbor]--
			if indegrees[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	if len(result) == n{
		mapValToIdx := make(map[int]int)
		for i, val := range result{
			mapValToIdx[val]=i
		}
		return mapValToIdx, true
	}
	return map[int]int{}, false
}

func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	rows, isRowOK := buildTopologicalSort(rowConditions, k)
	cols, isColOK := buildTopologicalSort(colConditions, k)
	if !isRowOK || !isColOK {
		return [][]int{}
	}
	result := make([][]int,k)
	for node := range result {
		result[node]=make([]int,k)
	}
	for i := 1; i <= k; i++ {
		result[rows[i]][cols[i]]=i
	}
	return result
}
