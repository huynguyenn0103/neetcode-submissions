func checkRow(row int, board [][]byte)bool{
	visited := make(map[byte]bool)
	for i := 0; i < 9; i++ {
		cell := board[row][i]
		if cell == '.' {continue}
		if visited[cell]{return false}
		visited[cell] = true
	}
	return true
}

func checkCol(col int, board [][]byte)bool{
	visited := make(map[byte]bool)
	for i := 0; i < 9; i++ {
		cell := board[i][col]
		if cell == '.'{continue}
		if visited[cell]{return false}
		visited[cell]=true
	}
	return true
}

func checkGroup(topLeft, bottomRight []int, board [][]byte)bool{
	startRow, startCol := topLeft[0], topLeft[1]
	endRow, endCol := bottomRight[0], bottomRight[1]
	visited := make(map[byte]bool)
	for i := startRow; i <= endRow; i++ {
		for j := startCol; j <= endCol; j++{
			cell := board[i][j]
			if cell == '.'{continue}
			if visited[cell]{return false}
			visited[cell]=true
		}
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if !checkCol(i, board) || !checkRow(i, board){
			return false
		}
	}
	for i := 0; i < 9; i = i + 3{
		for j := 0; j < 9; j = j + 3{
			if !checkGroup([]int{i,j},[]int{i+2,j+2},board){
				return false
			}
		}
	}
	return true
}
