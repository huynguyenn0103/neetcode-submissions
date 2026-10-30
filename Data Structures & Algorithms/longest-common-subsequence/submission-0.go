func longestCommonSubsequence(text1 string, text2 string) int {
    m,n := len(text1),len(text2)
	matrix := make([][]int, (m+1)*(n+1))
	for row := range matrix{
		matrix[row]=make([]int, n+1)
	}
	for i:=1; i <=m; i++{
		for j:=1; j<=n; j++{
			if text1[i-1] == text2[j-1]{
				tmp := 1
				tmp += matrix[i-1][j-1]
				matrix[i][j]=tmp
			} else{
				matrix[i][j]=max(matrix[i][j-1],matrix[i-1][j])
			}
		}
	}
	return matrix[m][n]
}
