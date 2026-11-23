var rowActions = []int{-1,0,1,0}
var colActions = []int{0,1,0,-1}

func longestIncreasingPath(matrix [][]int) int {
    cache := make(map[[2]int]int)
	m,n := len(matrix),len(matrix[0])
	var longestPathFrom func(row, col int)int
	longestPathFrom = func(row, col int)int{
		key := [2]int{row,col}
		if val, existed := cache[key]; existed{
			return val
		}
		result := 1
		for i := 0; i < len(rowActions); i++{
			nextRow := row + rowActions[i]
			nextCol := col + colActions[i]
			if nextRow >= 0 && nextRow < m && nextCol >= 0 && nextCol < n && matrix[nextRow][nextCol] > matrix[row][col]{
				result = max(result,1+longestPathFrom(nextRow,nextCol))
			}
		}
		cache[key]=result
		return result
	}
	result := 1
	for i := 0; i < m; i++{
		for j := 0; j < n; j++{
			result = max(longestPathFrom(i,j),result)
		}
	}
	return result
}
