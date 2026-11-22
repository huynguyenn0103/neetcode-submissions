const NotSurround = byte('N')

type Cell struct{
	row int
	col int
}

var rowActions = []int{-1,0,1,0}
var colActions = []int{0,1,0,-1}

func solve(board [][]byte) {
    m, n := len(board), len(board[0])
	var handleFromCell func(i,j int)
	handleFromCell = func(i,j int){
		if board[i][j]!='O'{return}
		queue := make([]Cell,0)
		queue = append(queue, Cell{row: i, col: j})
		for len(queue) > 0{
			cell := queue[0]
			queue = queue[1:]
			row, col := cell.row, cell.col
			if board[row][col] == NotSurround{continue}
			board[row][col]=NotSurround
			for i := 0; i < len(rowActions); i++{
				newRow := row + rowActions[i]
				newCol := col + colActions[i]
				if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n && board[newRow][newCol]=='O'{
					queue = append(queue, Cell{row: newRow, col: newCol})
				}
			}
		}
	}
	for i := 0; i < m; i++{
		handleFromCell(i,0)
		handleFromCell(i,n-1)
	}
	for j := 0; j < n; j++{
		handleFromCell(0,j)
		handleFromCell(m-1,j)
	}
	for i := 0; i < m; i++{
		for j := 0; j < n; j++{
			if board[i][j] == 'O'{
				board[i][j] = 'X'
			}
			if board[i][j] == NotSurround{
				board[i][j] = 'O'
			}
		}
	}
}
