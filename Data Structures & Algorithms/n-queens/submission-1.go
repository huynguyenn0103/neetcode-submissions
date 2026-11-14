import "slices"

type Pos struct{
	row int
	col int
}

func abs(x int)int{
	if x >= 0 {
		return x
	}
	return -x
}

func isValidBetweenTwoQueens(q1, q2 Pos)bool{
	row1, col1 := q1.row, q1.col
	row2, col2 := q2.row, q2.col
	if row1 == row2 || col1 == col2 || abs(row1-row2) == abs(col1-col2){
		return false
	}
	return true
}

func canAddedNewQueen(newQueen Pos, queens []Pos)bool{
	for _, q := range queens{
		if !isValidBetweenTwoQueens(newQueen, q){
			return false
		}
	} 
	return true
}

func nextPos(pos Pos, rows, cols int) (Pos,bool){
	row, col := pos.row, pos.col
	if col < cols - 1{
		return Pos{row: row, col: col + 1},true
	} else {
		newRow := row + 1
		if newRow >= rows {
			return Pos{},false
		}
		return Pos{row: newRow, col: 0},true
	}
}

func solveNQueens(n int) [][]string {
	results := make([][]Pos,0)
	combination := make([]Pos,0)
	var backtrack func(pos Pos)
	backtrack = func(pos Pos){
		for {
			next, isValid := nextPos(pos,n,n)
			if canAddedNewQueen(pos, combination){
				combination = append(combination, pos)
				if len(combination) == n{
					results = append(results, slices.Clone(combination))
				} else if isValid{
					backtrack(next)
				}
				combination = combination[:len(combination)-1]
			}
			if !isValid {break}
			pos = next
		}
	}
	backtrack(Pos{0,0})
	resultInStrs := make([][]string,0)
	for _, result := range results {
		layout := make([]string, 0)
		for i := 0; i < n; i++ {
			var builder strings.Builder
			for j := 0; j < n; j++{
				if len(result) > 0 && result[0].row == i && result[0].col == j{
					builder.WriteString("Q")
					result = result[1:]
				} else {
					builder.WriteString(".")
				}
			}
			layout = append(layout, builder.String())
		}
		resultInStrs = append(resultInStrs, layout)
	}
	return resultInStrs
}
