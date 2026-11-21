var rowActions = []int{-1,0,1,0}
var colActions = []int{0,1,0,-1}

func exist(board [][]byte, word string) bool {
	visited := make(map[[2]int]bool)
	l := len(word)
	m, n := len(board), len(board[0])
	var backtrack func(row, col, idx int)bool
	backtrack = func(row,col,idx int)bool{
		if board[row][col] != word[idx]{
			return false
		}
		if idx == l-1{
			return true
		}
		var result bool
		visited[[2]int{row,col}]=true
		for i := 0; i < len(rowActions); i++ {
			nextRow := row + rowActions[i]
			nextCol := col + colActions[i]
			if nextRow >= 0 && nextRow < m && nextCol >= 0 && nextCol < n && !visited[[2]int{nextRow, nextCol}]{
				result = backtrack(nextRow,nextCol,idx+1)
				if result{break}
			}
		}
		visited[[2]int{row,col}]=false
		return result
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backtrack(i,j,0){
				return true
			}
		}
	}
	return false
}
