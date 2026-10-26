const Infinity = 1<<31 - 1
type Point struct {
	row int
	col int
	level int
}

var rowActions = []int{-1,0,1,0}
var colActions = []int{0,1,0,-1}

func islandsAndTreasure(grid [][]int) {
	m, n := len(grid), len(grid[0])
			queue := make([]Point, 0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {continue}
			queue = append(queue, Point{row: i, col: j, level: 0})
		}
	}

	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]
		row, col, level := point.row, point.col, point.level
		nextLevel := level + 1
		for i := 0; i < len(rowActions); i++ {
			nextRow := row + rowActions[i]
			nextCol := col + colActions[i]
			if nextRow >= 0 && nextRow < m && nextCol >= 0 && nextCol < n && grid[nextRow][nextCol]==Infinity {
				queue = append(queue, Point{row: nextRow, col: nextCol, level: nextLevel})
				grid[nextRow][nextCol]=nextLevel
			}
		}
	}
}
