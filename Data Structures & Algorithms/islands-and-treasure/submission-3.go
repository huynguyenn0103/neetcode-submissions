const INF = (1<<31) - 1

var rowActions = []int{-1,0,1,0}
var colActions = []int{0,1,0,-1}

type Cell struct{
	row int
	col int
	level int
}

func islandsAndTreasure(grid [][]int) {
    queue := make([]Cell, 0)
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++{
		for j := 0; j < n; j++ {
			if grid[i][j] == 0{
				queue = append(queue, Cell{row: i, col: j, level: 0})
			}
		}
	}
	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		row,col,level := cell.row,cell.col,cell.level
		if grid[row][col] != INF && grid[row][col] != 0{continue}
		grid[row][col]=level
		for i := 0; i < len(rowActions); i++{
			newRow := row + rowActions[i]
			newCol := col + colActions[i]
			if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n && grid[newRow][newCol] == INF{
				queue = append(queue, Cell{row: newRow, col: newCol, level: level+1})
			}
		}
	}
}
