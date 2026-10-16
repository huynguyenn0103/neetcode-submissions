func isValidSudoku(board [][]byte) bool {
	n := len(board)
	for i := 0; i < n; i++ {
		visitedByRow := make(map[byte]bool)
		for j := 0; j < n; j++ {
			val := board[i][j]
			if visitedByRow[val] && val != '.'{
				return false
			}
			visitedByRow[val] = true
		}
	}

	for i := 0; i < n; i++ {
		visitedByCol := make(map[byte]bool)
		for j := 0; j < n; j++ {
			val := board[j][i]
			if visitedByCol[val] && val != '.'{
				return false
			}
			visitedByCol[val] = true
		}
	}

	for startRow := 0; startRow < n; startRow = startRow + 3{
		for startCol := 0; startCol < n; startCol = startCol + 3{
			visitedByBox := make(map[byte]bool)
			for i := startRow; i < startRow + 3; i++ {
				for j := startCol; j < startCol + 3; j++ {
					val := board[i][j]
					if visitedByBox[val] && val != '.'{
						return false
					}
					visitedByBox[val]=true
				}
			}
		}
	}
	return true
}
