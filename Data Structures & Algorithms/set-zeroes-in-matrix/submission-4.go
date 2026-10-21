func setZeroes(matrix [][]int) {
    firstCol, lastCol, firstRow, lastRow := false, false, false, false
	m, n := len(matrix), len(matrix[0])

	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			firstCol = true

		}
		if matrix[i][n-1] == 0 {
			lastCol = true
		}
	}
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			firstRow = true
		}
		if matrix[m-1][j] == 0 {
			lastRow = true
		}
	}
	
	for i := 1 ; i < m - 1; i++ {
		for j := 1; j < n - 1; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[m-1][j] = 0
				matrix[i][0] = 0
				matrix[i][n-1] = 0
			}
		}
	}
	for i := 1; i < m - 1; i++ {
		for j := 1; j < n - 1; j++ {
			if matrix[0][j] == 0 || matrix[m-1][j] == 0 || matrix[i][0] == 0 || matrix[i][n-1] == 0{
				matrix[i][j] = 0
				if matrix[i][0] == 0 || matrix[i][n-1] == 0 {
					matrix[i][0] = 0
					matrix[i][n-1] = 0
				}
				if matrix[0][j] == 0 || matrix[m-1][j] == 0 {
					matrix[0][j] = 0
					matrix[m-1][j] = 0
				}
			}
		}
	}
	if firstCol {
		for i := 0; i < m; i++{
			matrix[i][0]=0
		}
	}
	if lastCol {
		for i := 0; i < m; i++{
			matrix[i][n-1] = 0
		}
	}
	if firstRow{
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
	if lastRow {
		for j := 0; j < n; j++ {
			matrix[m-1][j] = 0
		}
	}
}
