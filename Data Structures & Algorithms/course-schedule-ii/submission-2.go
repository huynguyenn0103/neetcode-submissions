func findOrder(numCourses int, prerequisites [][]int) []int {
    graph := make([][]int, numCourses)
	indegrees := make([]int,numCourses)
	for _, prerequisite := range prerequisites{
		from, to := prerequisite[1], prerequisite[0]
		graph[from]=append(graph[from],to)
		indegrees[to]++
	}
	result := make([]int,0)
	queue := make([]int,0)
	for node, degree := range indegrees{
		if degree == 0{
			queue = append(queue, node)
		}
	}
	for len(queue) > 0{
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)
		for _, neighborNode := range graph[node]{
			indegrees[neighborNode]--
			if indegrees[neighborNode] == 0{
				queue = append(queue, neighborNode)
			}
		}
	}
	if len(result) == numCourses{
		return result
	}
	return []int{}
}
