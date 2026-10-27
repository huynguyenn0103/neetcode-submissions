type NodeItem struct {
	elevation int
	row int
	col int
}

var rowActions = []int{-1,0,1,0}
var colActions = []int{0,1,0,-1}

type MinHeap []NodeItem
func (h MinHeap) Len()int{return len(h)}
func (h MinHeap) Less(i,j int)bool{return h[i].elevation<h[j].elevation}
func (h MinHeap) Swap(i,j int){h[i],h[j]=h[j],h[i]}
func (h *MinHeap) Push(x interface{}){*h = append(*h, x.(NodeItem))}
func (h *MinHeap) Pop()interface{}{
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

const Infinity = math.MaxInt

func swimInWater(grid [][]int) int {
	n := len(grid)
	result := grid[0][0]
    minHeap := &MinHeap{}
	heap.Init(minHeap)
	heap.Push(minHeap, NodeItem{row: 0, col: 0, elevation: grid[0][0]})
	for minHeap.Len() > 0 {
		item := heap.Pop(minHeap).(NodeItem)
		row, col, elevation := item.row, item.col, item.elevation
		if grid[row][col] == Infinity{continue}
		grid[row][col]=Infinity
		result = max(result, elevation)
		if row == n-1 && col == n-1{
			return result
		}
		for i := 0; i < len(rowActions); i++{
			nextRow := row + rowActions[i]
			nextCol := col + colActions[i]
			if nextRow >= 0 && nextRow < n && nextCol >= 0 && nextCol < n && grid[nextRow][nextCol] != Infinity{
				heap.Push(minHeap, NodeItem{row: nextRow, col: nextCol, elevation: grid[nextRow][nextCol]})
			}
		}
	}
	return result
}
