func findOrder(numCourses int, prerequisites [][]int) []int {
	result := make([]int, 0)
    graph := make([][]int, numCourses)
	for node := range graph{
		graph[node] = make([]int, 0)
	}
	indegrees := make([]int, numCourses)
	for _, prerequisite := range prerequisites{
		from, to := prerequisite[1], prerequisite[0]
		graph[from] = append(graph[from], to)
		indegrees[to]++
	}
	queue := make([]int, 0)
	for node, indegree := range indegrees{
		if indegree == 0 {
			queue = append(queue, node)
		}
	}
	processCourses := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		processCourses++
		result = append(result, node)
		for _, neighbor := range graph[node]{
			indegrees[neighbor]--
			if indegrees[neighbor] == 0{
				queue = append(queue, neighbor)
			}
		}
	}
	if processCourses == numCourses {
		return result
	}
	return []int{}
}
