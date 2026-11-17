import "slices"

const INF = math.MaxInt

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
    dist := make([]int, n)
	for node := range dist{
		dist[node]=INF
	}
	dist[src]=0
	for turn := 0; turn <= k; turn++{
		clone := slices.Clone(dist)
		for _, flight := range flights{
			from, to, weight := flight[0],flight[1],flight[2]
			if dist[from] == INF{continue}
			if dist[from] + weight < clone[to]{
				clone[to] = dist[from]+weight
			}
		}
		dist = clone
	}
	if dist[dst] == INF{
		return -1
	}
	return dist[dst]
}
