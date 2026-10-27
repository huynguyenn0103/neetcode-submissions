import "slices"
const Infinity = math.MaxInt
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
   dist := make([]int, n) 
   for node := range dist{
	dist[node]=Infinity
   }
   dist[src]=0
   for i := 0; i <= k; i++ {
	 cloneDist := slices.Clone(dist)
	 for _, edge := range flights {
		from, to, weight := edge[0],edge[1],edge[2]
		if cloneDist[from] == Infinity{continue}
		if cloneDist[from] + weight < dist[to]{
			dist[to] = cloneDist[from] + weight
		}
	 }
   }
   if dist[dst] == Infinity{return -1}
   return dist[dst]
}
