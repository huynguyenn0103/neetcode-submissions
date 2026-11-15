type Cell struct{
	row int
	col int
	minTimeFromStart int
}

type MinHeap []Cell
func (h MinHeap) Len()int{return len(h)}
func (h MinHeap) Less(i,j int)bool{return h[i].minTimeFromStart < h[j].minTimeFromStart}
func (h MinHeap) Swap(i,j int){h[i],h[j] = h[j],h[i]}
func (h *MinHeap) Push(x interface{}){*h = append(*h, x.(Cell))}
func (h *MinHeap) Pop()interface{}{
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

var rowActions = []int{-1,0,1,0}
var colActions = []int{0,1,0,-1}

const INF = math.MaxInt

func swimInWater(grid [][]int) int {
    m,n := len(grid), len(grid[0])
	minTimes := make([][]int,m)
	for i := 0; i < m; i++ {
		minTimes[i] = make([]int,n)
	}
	for i := 0; i < m; i++{
		for j := 0; j < n; j++{
			minTimes[i][j]=INF
		}
	}
	minTimes[0][0]=grid[0][0]
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	heap.Push(minHeap, Cell{row: 0, col: 0, minTimeFromStart: grid[0][0]})
	for minHeap.Len() > 0 {
		cell := heap.Pop(minHeap).(Cell)
		row, col, minTimeFromStart := cell.row, cell.col, cell.minTimeFromStart
		if minTimes[row][col] != minTimeFromStart{continue}
		for i := 0; i < len(rowActions); i++{
			newRow := row + rowActions[i]
			newCol := col + colActions[i]
			if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n{
				newTime := max(minTimeFromStart, grid[newRow][newCol])
				if minTimes[newRow][newCol] > newTime{
					minTimes[newRow][newCol] = newTime
					heap.Push(minHeap, Cell{row: newRow, col: newCol, minTimeFromStart: newTime})
				}
			}
		}
	}
	return minTimes[m-1][n-1]
}
