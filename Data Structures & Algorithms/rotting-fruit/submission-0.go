const Visited = 3

type Point struct {
	row int
	col int
	level int
}
var rowActions = []int{-1,0,1,0}
var colActions = []int{0,1,0,-1}

func orangesRotting(grid [][]int) int {
    numOfFruit := 0
	numOfRotten := 0
	m, n := len(grid), len(grid[0])
	queue := make([]Point, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++{
			if grid[i][j] != 0 {
				numOfFruit++
				if grid[i][j] == 2{
					queue = append(queue, Point{row: i, col: j, level: 0})
				}
			}
		}
	}
	time := 0
	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]
		row, col, level := point.row, point.col, point.level
		if grid[row][col] == Visited{continue}
		numOfRotten++
		time = max(time, level)
		grid[row][col]=Visited
		for i := 0; i < len(rowActions);i++{
			nextRow := row + rowActions[i]
			nextCol := col + colActions[i]
			if nextRow >= 0 && nextRow < m && nextCol >= 0 && nextCol < n && grid[nextRow][nextCol] !=0{
				queue = append(queue, Point{row: nextRow, col: nextCol, level: level + 1})
			}
		}
	}
	if numOfRotten == numOfFruit {
		return time
	}
	return -1
}
