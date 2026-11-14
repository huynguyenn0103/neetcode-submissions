func canFinish(numCourses int, prerequisites [][]int) bool {
    graph := make([][]int, numCourses)
	for node := range graph{
		graph[node]=make([]int,0)
	}
	indegrees := make([]int, numCourses)
	for _, prerequisite := range prerequisites{
		from, to := prerequisite[1], prerequisite[0]
		graph[from]=append(graph[from],to)
		indegrees[to]++
	}
	queue := make([]int,0)
	for node, indegree := range indegrees{
		if indegree == 0{
			queue = append(queue,node)
		}
	}
	numOfProcessNode := 0
	for len(queue) > 0{
		node := queue[0]
		queue = queue[1:]
		numOfProcessNode++
		for _, neighborNode := range graph[node]{
			indegrees[neighborNode]--
			if indegrees[neighborNode] == 0{
				queue = append(queue, neighborNode)
			}
		}
	}
	return numOfProcessNode == numCourses
}
