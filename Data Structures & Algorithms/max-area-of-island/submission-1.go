const Visited = 2
var rowActions = []int{-1,0,1,0}
var colActions = []int{0,1,0,-1}

type Cell struct {
	row int
	col int
}

func maxAreaOfIsland(grid [][]int) int {
    m, n := len(grid), len(grid[0])
	maxArea := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == Visited || grid[i][j] == 0 {continue}
			queue := make([]Cell, 0)
			queue = append(queue, Cell{row: i, col: j})
			area := 0
			for len(queue) > 0 {
				cell := queue[0]
				queue = queue[1:]
				row, col := cell.row, cell.col
				if grid[row][col] == Visited{
					continue
				}
				grid[row][col] = Visited
				area++
				for k := 0; k < len(rowActions); k++{
					newRow := row + rowActions[k]
					newCol := col + colActions[k]
					if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n && grid[newRow][newCol] == 1{
						queue = append(queue, Cell{row: newRow, col: newCol})
					}
				}
			}
			maxArea = max(maxArea, area)
		}
	}
	return maxArea
}
