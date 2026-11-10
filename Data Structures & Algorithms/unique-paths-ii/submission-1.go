
var rowActions = []int{0,1}
var colActions = []int{1,0}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	cache := make(map[[2]int]int)
	var uniquePathFromPoint func(row, col int)int
	uniquePathFromPoint = func(row, col int)int{
		if obstacleGrid[row][col] == 1 {
			return 0
		}
		if row == m-1 && col == n-1 {
			if obstacleGrid[row][col] == 0{
				return 1
			}
			return 0
		}
		if val, existed := cache[[2]int{row,col}]; existed {
			return val
		}
		result := 0
		for i := 0; i < len(rowActions); i++ {
			newRow := row + rowActions[i]
			newCol := col + colActions[i]
			if newRow >= 0 && newCol >= 0 && newRow < m && newCol < n && obstacleGrid[newRow][newCol] == 0 {
				result += uniquePathFromPoint(newRow,newCol)
			}
		}
		cache[[2]int{row,col}]=result
		return result
	}
	return uniquePathFromPoint(0,0)
}
