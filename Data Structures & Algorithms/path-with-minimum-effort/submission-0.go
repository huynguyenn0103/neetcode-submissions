type NodeItem struct {
	effortFromStart int
	row int
	col int
}

type MinHeap []NodeItem
func (h MinHeap)Len()int{return len(h)}
func (h MinHeap)Less(i,j int)bool{return h[i].effortFromStart < h[j].effortFromStart}
func (h MinHeap)Swap(i,j int){h[i],h[j]=h[j],h[i]}
func (h *MinHeap)Push(x interface{}){*h = append(*h, x.(NodeItem))}
func (h *MinHeap)Pop()interface{}{
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

var rowActions = []int{-1,0,1,0}
var colActions = []int{0,1,0,-1}

const Infinity = math.MaxInt

func abs(x int)int{
	if x > 0 {return x}
	return -x
}

func minimumEffortPath(heights [][]int) int {
	m,n := len(heights),len(heights[0])
	efforts := make([][]int,m)
	for i := 0; i < m; i++ {
		efforts[i]=make([]int,n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			efforts[i][j]=Infinity
		}
	}
	efforts[0][0]=0

	minHeap := &MinHeap{}
	heap.Init(minHeap)
	heap.Push(minHeap, NodeItem{
		effortFromStart: 0,
		row: 0,
		col: 0,
	})
	for minHeap.Len()>0{
		item := heap.Pop(minHeap).(NodeItem)
		row,col,effortFromStart := item.row,item.col,item.effortFromStart
		if efforts[row][col] != effortFromStart {continue}
		if row == m-1 && col == n-1{
			return effortFromStart
		}
		for i := 0; i < len(rowActions); i++ {
			newRow := row + rowActions[i]
			newCol := col + colActions[i]
			if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n {
				newEffort := max(effortFromStart, abs(heights[newRow][newCol]-heights[row][col]))
				if newEffort < efforts[newRow][newCol]{
					efforts[newRow][newCol]=newEffort
					heap.Push(minHeap, NodeItem{
						effortFromStart: newEffort,
						row: newRow,
						col: newCol,
					})
				}
			}
		}
	}
	return 0
}
