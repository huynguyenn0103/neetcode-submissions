const Visited = byte('2')

var rowActions = []int{-1,0,1,0}
var colActions = []int{0,1,0,-1}

type Cell struct{
	row int
	col int
}

func numIslands(grid [][]byte) int {
    numOfIsland := 0
	m,n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != '1' {continue}
			numOfIsland++
			queue := make([]Cell,0)
			queue = append(queue, Cell{row: i, col: j})
			for len(queue) > 0 {
				item := queue[0]
				queue = queue[1:]
				row, col := item.row, item.col
				if grid[row][col] == Visited {continue}
				grid[row][col] = Visited
				for i := 0; i < len(rowActions); i++ {
					newRow := row + rowActions[i]
					newCol := col + colActions[i]
					if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n && grid[newRow][newCol] == '1'{
						queue = append(queue, Cell{row: newRow, col: newCol})
					}
				}
			}
		}
	}
	return numOfIsland
}
