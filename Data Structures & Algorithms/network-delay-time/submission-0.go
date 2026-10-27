type NodeItem struct {
	node int
	distFromStart int
} 
type Node struct {
	node int
	weight int
}

type MinHeap []NodeItem
func(h MinHeap) Len()int{return len(h)}
func(h MinHeap) Less(i, j int)bool{return h[i].distFromStart < h[j].distFromStart}
func (h MinHeap) Swap(i, j int){h[i],h[j]=h[j],h[i]}
func (h *MinHeap) Push(x interface{}){*h = append(*h, x.(NodeItem))}
func (h *MinHeap) Pop()interface{}{
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

const Infinity = math.MaxInt

func networkDelayTime(times [][]int, n int, k int) int {
    graph := make([][]Node, n)
	for node := range graph{
		graph[node]=make([]Node, 0)
	}
	for _, time := range times{
		from, to, weight := time[0]-1, time[1]-1, time[2]
		graph[from]=append(graph[from], Node{node: to, weight: weight})
	}
	dist := make([]int, n)
	for node := range dist{
		dist[node]=Infinity
	}
	start := k-1
	dist[start]=0

	minHeap := &MinHeap{}
	heap.Init(minHeap)
	heap.Push(minHeap, NodeItem{node: start, distFromStart: 0})
	for minHeap.Len() > 0 {
		item := heap.Pop(minHeap).(NodeItem)
		node, distFromStart := item.node, item.distFromStart
		if dist[node] != distFromStart{continue}
		for _, neighbor := range graph[node]{
			newDist := distFromStart + neighbor.weight
			if newDist < dist[neighbor.node]{
				dist[neighbor.node]=newDist
				heap.Push(minHeap, NodeItem{node: neighbor.node, distFromStart: newDist})
			}
		}
	}

	result := -1
	for _, d := range dist{
		if d == Infinity{return -1}
		result = max(result, d)
	}
	return result
}
