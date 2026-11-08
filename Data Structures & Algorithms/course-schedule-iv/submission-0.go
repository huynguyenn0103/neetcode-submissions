func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
    graph := make([][]int, numCourses)

    for _, p := range prerequisites {
        a, b := p[0], p[1]
        graph[a] = append(graph[a], b)
    }

    reachable := make([][]bool, numCourses)

    for i := 0; i < numCourses; i++ {
        reachable[i] = make([]bool, numCourses)

        visited := make([]bool, numCourses)

        var dfs func(int)
        dfs = func(node int) {
            for _, next := range graph[node] {
                if !visited[next] {
                    visited[next] = true
                    reachable[i][next] = true
                    dfs(next)
                }
            }
        }

        dfs(i)
    }

    answer := make([]bool, len(queries))

    for i, q := range queries {
        answer[i] = reachable[q[0]][q[1]]
    }

    return answer
}