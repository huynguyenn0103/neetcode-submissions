type NumMatrix struct {
	sumMatrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	sumMatrix := make([][]int,m)
	for i := range sumMatrix{
		sumMatrix[i]=make([]int,n)
	}
	for i := 0; i < m; i++ {
		sum := 0
		for j := 0; j < n; j++ {
			sum += matrix[i][j]
			tmp := sum
			if i >= 1{
				tmp += sumMatrix[i-1][j]
			}
			sumMatrix[i][j]=tmp
		}
	}
	return NumMatrix{
		sumMatrix: sumMatrix,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	result := this.sumMatrix[row2][col2]
	if row1 >= 1{
		result = result - this.sumMatrix[row1-1][col2]
	}
	if col1 >= 1{
		result = result - this.sumMatrix[row2][col1-1]
	}
	if row1 >= 1 && col1 >= 1{
		result = result + this.sumMatrix[row1-1][col1-1]
	}
	return result
}

// Your NumMatrix object will be instantiated and called as such:
// obj := Constructor(matrix)
// param_1 := obj.SumRegion(row1,col1,row2,col2)
