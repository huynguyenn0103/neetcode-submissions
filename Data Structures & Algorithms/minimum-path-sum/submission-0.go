var rowActions = []int{0,1}
var colActions = []int{1,0}


func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	cache := make(map[[2]int]int)
	var minPathSumFrom func(row,col int)int
	minPathSumFrom = func(row,col int)int{
		if row == m-1 && col == n-1 {
			return grid[row][col]
		}
		if val,existed := cache[[2]int{row,col}]; existed{
			return val
		}
		result := math.MaxInt
		for i := 0; i < len(rowActions); i++ {
			newRow := row + rowActions[i]
			newCol := col + colActions[i]
			if newRow >= 0 && newCol >= 0 && newRow < m && newCol < n {
				result = min(result, grid[row][col] + minPathSumFrom(newRow,newCol))
			}
		}
		cache[[2]int{row,col}]=result
		return result
	}
	return minPathSumFrom(0,0)
}
