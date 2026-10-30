func uniquePaths(m int, n int) int {
    matrix := make([][]int,m*n)
	for row := range matrix{
		matrix[row]=make([]int,n)
	}
	matrix[0][0]=1
	for i:=0; i<m; i++{
		for j:=0; j<n; j++{
			if i==0 && j== 0{continue}
			valLeft, valTop := 0,0
			leftRow, leftCol := i, j-1
			if leftCol >= 0{
				valLeft = matrix[leftRow][leftCol]
			}
			topRow, topCol := i-1,j
			if topRow >= 0{
				valTop = matrix[topRow][topCol]
			}
            matrix[i][j]=valLeft + valTop
		}
	}
	return matrix[m-1][n-1]
}
