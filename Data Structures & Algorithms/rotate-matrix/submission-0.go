func rotate(matrix [][]int)  {
    n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i; j < n - 1 - i; j++ {
			k := 4
			curI, curJ := i, j
			prev := matrix[n-1-curJ][curI]
			for k != 0 {
				tmp := matrix[curI][curJ]
				matrix[curI][curJ] = prev
				prev = tmp
				curI, curJ = curJ, n-1-curI
				k--
			}
		}
	}
}
