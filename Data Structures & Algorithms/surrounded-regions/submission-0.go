const Visited = 'N'
type Point struct {
	row int
	col int
}

var rowActions = []int{-1,0,1,0}
var colActions = []int{0,1,0,-1}

func solve(board [][]byte) {
    m, n := len(board), len(board[0])
	queue := make([]Point, 0)
	for i := 0; i < m; i++{
		if board[i][0] == 'O' {
			queue = append(queue, Point{row: i, col: 0})
		}
		if board[i][n-1] == 'O' {
			queue = append(queue, Point{row: i, col: n-1})
		}
	}
	for j := 0; j < n; j++ {
		if board[0][j] == 'O'{
			queue = append(queue, Point{row: 0, col: j})
		}
		if board[m-1][j] == 'O'{
			queue = append(queue, Point{row: m-1, col: j})
		}
	}
	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]
		row, col := point.row, point.col
		if board[row][col] == Visited{continue}
		board[row][col]=Visited
		for i := 0; i < len(rowActions); i++ {
			nextRow := row + rowActions[i]
			nextCol := col + colActions[i]
			if nextRow >= 0 && nextRow < m && nextCol >= 0 && nextCol < n && board[nextRow][nextCol] == 'O'{
				queue = append(queue, Point{row: nextRow, col: nextCol})
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O'{
				board[i][j] = 'X'
			} else if board[i][j] == Visited{
				board[i][j] = 'O'
			}
		}
	}
}
