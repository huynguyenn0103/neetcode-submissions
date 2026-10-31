var rowActions = []int{-1,0,1,0}
var colActions = []int{0,1,0,-1}

func longestIncreasingPath(matrix [][]int) int {
	m, n := len(matrix),len(matrix[0])
    cache := make(map[[2]int]int)
	var longestFromPoint func(row, col int)int
	longestFromPoint = func(row, col int)int{
		key := [2]int{row,col}
		if value, existed := cache[key]; existed{
			return value
		}
		longest := 0
		for i := 0; i < len(rowActions); i++{
			nextRow := row + rowActions[i]
			nextCol := col + colActions[i]
			if nextRow >= 0 && nextRow < m && nextCol >= 0 && nextCol < n {
				if matrix[nextRow][nextCol] > matrix[row][col]{
					longest = max(longest, longestFromPoint(nextRow,nextCol))
				}
			}
		}
		cache[key]=longest+1
		return cache[key]
	}
	result := 0
	for i:=0; i<m; i++ {
		for j:=0; j<n; j++{
			result = max(result, longestFromPoint(i,j))
		}
	}
	return result
}
