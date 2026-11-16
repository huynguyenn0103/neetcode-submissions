type Node struct{
	node int
	weight int
}

type NodeItem struct{
	timeFromStart int
	node int
}

type MinHeap []NodeItem
func (h MinHeap) Len()int{return len(h)}
func (h MinHeap) Less(i,j int)bool{return h[i].timeFromStart < h[j].timeFromStart}
func (h MinHeap) Swap(i,j int){h[i],h[j]=h[j],h[i]}
func (h *MinHeap) Push(x interface{}){*h = append(*h, x.(NodeItem))}
func (h *MinHeap) Pop()interface{}{
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

const INF = math.MaxInt

func networkDelayTime(times [][]int, n int, k int) int {
    graph := make([][]Node, n)
	for node := range graph{
		graph[node] = make([]Node,0)
	}
	for _, time := range times{
		from,to,weight := time[0]-1, time[1]-1, time[2]
		graph[from] = append(graph[from], Node{node: to, weight: weight})
	}
	timesToNode := make([]int, n)
	for node := range timesToNode{
		timesToNode[node]=INF
	}
	timesToNode[k-1]=0

	minHeap := &MinHeap{}
	heap.Init(minHeap)
	heap.Push(minHeap, NodeItem{node: k-1, timeFromStart: 0})
	
	for minHeap.Len() > 0 {
		item := heap.Pop(minHeap).(NodeItem)
		node, timeFromStart := item.node, item.timeFromStart
		if timesToNode[node] != timeFromStart{continue}
		for _, neighbor := range graph[node]{
			neighborNode, weight := neighbor.node, neighbor.weight
			newTime := timeFromStart + weight
			if newTime < timesToNode[neighborNode]{
				timesToNode[neighborNode]=newTime
				heap.Push(minHeap, NodeItem{node: neighborNode, timeFromStart: newTime})
			}
		}
	}

	maxTime := 0
	for _, time := range timesToNode{
		if time == INF{return -1}
		maxTime = max(maxTime, time)
	}
	return maxTime
}
