func findMinHeightTrees(n int, edges [][]int) []int {
    if n == 1 {
        return []int{0}
    }

    // 1. Build adjacency list
    graph := make([][]int, n)

    for _, e := range edges {
        a, b := e[0], e[1]
        graph[a] = append(graph[a], b)
        graph[b] = append(graph[b], a)
    }

    // 2. Calculate degree of each node
    degree := make([]int, n)

    leaves := []int{}

    for i := 0; i < n; i++ {
        degree[i] = len(graph[i])

        if degree[i] == 1 {
            leaves = append(leaves, i)
        }
    }

    remaining := n

    // 3. Remove leaves layer by layer
    for remaining > 2 {
        size := len(leaves)
        remaining -= size

        nextLeaves := []int{}

        for _, leaf := range leaves {
            for _, neighbor := range graph[leaf] {
                degree[neighbor]--

                if degree[neighbor] == 1 {
                    nextLeaves = append(nextLeaves, neighbor)
                }
            }
        }

        leaves = nextLeaves
    }

    return leaves
}