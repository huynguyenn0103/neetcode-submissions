func transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	result := make([][]int, n)
	for row := range result {
		result[row]=make([]int,m)
	}
	for row, cols := range matrix{
		for col := range cols {
			result[col][row]=matrix[row][col]
		}
	}
	return result
}
