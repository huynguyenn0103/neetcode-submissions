func uniquePaths(m int, n int) int {
	cacheNumOfPathFromPos := make(map[[2]int]int)
	var numOfUniquePathFromPos func(row,col int)int
	numOfUniquePathFromPos = func(row, col int)int{
		if row == m-1 && col == n-1{
			return 1
		}
		key := [2]int{row,col}
		if val, existed := cacheNumOfPathFromPos[key]; existed{
			return val
		}
		result := 0
		if col < n-1 {
			result += numOfUniquePathFromPos(row, col+1)
		}
		if row < m-1 {
			result += numOfUniquePathFromPos(row+1, col)
		}
		cacheNumOfPathFromPos[key]=result
		return result
	}
	return numOfUniquePathFromPos(0,0)
}
