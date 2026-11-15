type RottenCell struct{
	row int
	col int
	level int
}

const Visited = 3
var rowActions = []int{-1,0,1,0}
var colActions = []int{0,1,0,-1}

func orangesRotting(grid [][]int) int {
    numOfFruit := 0
	queue := make([]RottenCell, 0)
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			c := grid[i][j]
			if c == 0 {continue}
			if c == 2 {
				queue = append(queue, RottenCell{row: i, col: j, level: 0})
			}
			numOfFruit++
		}
	}
	numOfRotten := 0
	result := 0
	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		row,col,level := cell.row, cell.col, cell.level
		if grid[row][col] == Visited{continue}
		grid[row][col] = Visited
		result = max(result, level)
		numOfRotten++
		for i := 0; i < len(rowActions); i++{
			newRow := row + rowActions[i]
			newCol := col + colActions[i]
			if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n && grid[newRow][newCol] == 1{
				grid[newRow][newCol] = 1
				queue = append(queue, RottenCell{row: newRow, col: newCol, level: level+1})
			}
		}
	}
	if numOfRotten == numOfFruit{
		return result
	}
	return -1
}
