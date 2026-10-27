

func findItinerary(tickets [][]string) []string {
    graph := make(map[string][]string)

    // Build adjacency list
    for _, ticket := range tickets {
        from, to := ticket[0], ticket[1]
        graph[from] = append(graph[from], to)
    }

    // Sort destinations in reverse lexicographical order
    for airport := range graph {
        sort.Sort(sort.Reverse(sort.StringSlice(graph[airport])))
    }

    result := make([]string, 0, len(tickets)+1)

    var dfs func(string)

    dfs = func(airport string) {
        for len(graph[airport]) > 0 {
            n := len(graph[airport])

            // Get smallest destination
            next := graph[airport][n-1]

            // Remove the ticket
            graph[airport] = graph[airport][:n-1]

            dfs(next)
        }

        // Add airport after using all outgoing tickets
        result = append(result, airport)
    }

    dfs("JFK")

    // Reverse result
    for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
        result[i], result[j] = result[j], result[i]
    }

    return result
}