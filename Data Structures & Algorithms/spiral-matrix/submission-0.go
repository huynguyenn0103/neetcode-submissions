func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	result := make([]int,0)
    startI, startJ := 0, 0
	for len(result) != m * n {
       curI, curJ := startI, startJ
	   for curJ <= n - 1 - startJ {
		result = append(result,matrix[curI][curJ])
		curJ++
	   }
	   curJ--
	   curI++
	   if len(result) == m*n {
		 break;
	   }
	   for curI <= m - 1 - startI {
		result = append(result, matrix[curI][curJ])
		curI++
	   }
	   curI--
	   curJ--
	   if len(result)==m*n{
		break
	   }
	   for curJ >= startJ {
		result = append(result,matrix[curI][curJ])
		curJ--
	   }
	   curJ++
	   curI--
	   if len(result)==m*n{
		break
	   }
	   for curI > startI {
		result = append(result, matrix[curI][curJ])
		curI--
	   }
	   startI = startI + 1
	   startJ = startJ + 1
	}
	return result
}
