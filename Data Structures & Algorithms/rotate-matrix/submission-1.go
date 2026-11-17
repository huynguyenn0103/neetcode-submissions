//[3,3] 0,0 -> 2,0 | 2,1 -> 1,2 | 2,0 -> 0,2 | 1,0 -> 0,1

func rotate(matrix [][]int)  {
   n := len(matrix)
   newMatrix := make([][]int,n)
   for row := range newMatrix{
	newMatrix[row]=make([]int,n)
   } 
   for i := 0; i < n; i++ {
	   for j := 0; j < n; j++ {
		newMatrix[i][j] = matrix[n-1-j][i]
	   }
   }
   for i := 0; i < n; i++ {
	for j := 0; j < n; j++ {
		matrix[i][j] = newMatrix[i][j]
	}
   }
}
