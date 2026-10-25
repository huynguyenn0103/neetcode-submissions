var rowActions = []int{-1,0,1,0}
var colActions = []int{0,1,0,-1}

func exist(board [][]byte, word string) bool {
	m,n := len(board), len(board[0])
	var backtrack func(row, col int, word string)bool
	backtrack = func(row, col int, word string)bool{
		c := word[0]
		if len(word) <= 1{
			return c == board[row][col]
		}
		if board[row][col] != c{
			return false
		}
		board[row][col] = '0'
		for i := 0; i < len(rowActions); i++{
			newRow := row + rowActions[i]
			newCol := col + colActions[i]
			if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n{
				if board[newRow][newCol] == '0' {continue}
				isSuccess := backtrack(newRow, newCol, word[1:])
				if isSuccess {
					return true
				}
			}
		}
		board[row][col]=c
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backtrack(i,j,word){
				return true
			}
		}
	}
	return false
}
