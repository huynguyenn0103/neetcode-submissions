var rowActions = []int{-1,0,1,0}
var colActions = []int{0,1,0,-1}
const Visited = 2

type Point struct {
	row int
	col int
}

func maxAreaOfIsland(grid [][]int) int {
    m, n := len(grid), len(grid[0])
	result := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == Visited || grid[i][j] == 0 {continue}
			area := 0
			queue := make([]Point, 0)
			queue = append(queue, Point{row: i, col: j})
			for len(queue) > 0 {
				point := queue[0]
				queue = queue[1:]
				row, col := point.row, point.col
				if grid[row][col]==Visited{continue}
				grid[row][col]=Visited
				area++
				for i := 0 ; i < len(rowActions); i++{
					nextRow := row + rowActions[i]
					nextCol := col + colActions[i]
					if nextRow >= 0 && nextRow < m && nextCol >= 0 && nextCol < n && grid[nextRow][nextCol]==1 {
						queue = append(queue, Point{row: nextRow, col: nextCol})
					}
				}
			}
			result = max(result, area)
		}
	}
	return result
}
