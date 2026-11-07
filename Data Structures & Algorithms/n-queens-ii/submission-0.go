func abs(x int)int {
	if x < 0 {
		return -x
	}
	return x
}

func isValid(newQueen []int, existedQueens [][]int) bool{
	newRow,newCol := newQueen[0], newQueen[1]
	for _, existQueen := range existedQueens{
		existRow,existCol := existQueen[0],existQueen[1]
		if existRow == newRow || existCol == newCol || abs(existRow-newRow) == abs(existCol-newCol){
			return false
		}
	}
	return true
}

func nextPos(curPos []int, n int)[]int{
	curRow, curCol := curPos[0], curPos[1]
	if curCol == n-1 {
		return []int{curRow+1,0}
	}
	return []int{curRow,curCol+1}
}

func totalNQueens(n int) int {
	existedQueens := make([][]int,0)
	result := 0
	var backtrack func(curPos []int)
	backtrack = func(curPos []int){
		if len(existedQueens) == n {
			result++
			return
		}
		curRow, curCol := curPos[0], curPos[1]
		if curRow >= n {
			return
		}
		for curRow < n {
			for curCol < n {
				newQueen := []int{curRow,curCol}
				if isValid(newQueen, existedQueens){
					existedQueens = append(existedQueens, newQueen)
					backtrack(nextPos(newQueen,n))
					existedQueens = existedQueens[:len(existedQueens)-1]
				}
				curCol++
			}
			curCol = 0
			curRow++
		}
	}
	backtrack([]int{0,0})
	return result
}
