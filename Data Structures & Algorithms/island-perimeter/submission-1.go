var rowActions = []int{-1,0,1,0}
var colActions = []int{0,1,0,-1}

func islandPerimeter(grid [][]int) int {
	m,n := len(grid), len(grid[0])
	result := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1{
				total := 4
				for k := 0; k < len(rowActions); k++ {
					newRow := i + rowActions[k]
					newCol := j + colActions[k]
					if newRow >= 0 && newCol >= 0 && newRow < m && newCol < n && grid[newRow][newCol] == 1 {
						total--
					}
				}
				result = result + total
			}
		}
	}
	return result
}
