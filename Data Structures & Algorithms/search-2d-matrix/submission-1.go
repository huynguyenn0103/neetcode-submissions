// 11 ~ [2,3] rows = 3 & cols = 4 => row * cols + col

func findMidCell(left, right [2]int, cols int)[2]int{
	leftCellToNum := left[0] * cols + left[1]
	rightCellToNum := right[0] * cols + right[1]
	mid := (leftCellToNum + rightCellToNum) / 2
	return [2]int{mid / cols, mid - mid/cols * cols}
}

func nextCell(curCell [2]int, rows, cols int)([2]int,bool){
	curRow, curCol := curCell[0], curCell[1]
	curCol = curCol + 1
	if curCol >= cols{
		curRow = curRow + 1
		curCol = 0
	}
	if curRow < rows && curCol < cols{
		return [2]int{curRow,curCol}, true
	}
	return [2]int{}, false
}

func prevCell(curCell [2]int, rows, cols int)([2]int,bool){
	curRow, curCol := curCell[0], curCell[1]
	curCol = curCol - 1
	if curCol < 0{
		curRow = curRow - 1
		curCol = cols - 1
	}
	if curRow >= 0 && curCol >= 0{
		return [2]int{curRow,curCol}, true
	}
	return [2]int{}, false
}

func isLeftSmallerOrEqual(left, right [2]int)bool{
	if left[0] == right[0]{
		return left[1] <= right[1]
	}
	return left[0] < right[0]
}

func searchMatrix(matrix [][]int, target int) bool {
	m,n := len(matrix), len(matrix[0])
	left := [2]int{0,0}
	right := [2]int{m-1,n-1}
	for isLeftSmallerOrEqual(left,right){
		mid := findMidCell(left,right,n)
		if matrix[mid[0]][mid[1]] == target{
			return true
		} else if matrix[mid[0]][mid[1]] > target{
			var isOk bool
			right, isOk = prevCell(right, m, n)
			if !isOk {
				break
			}
		} else {
			var isOk bool
			left, isOk = nextCell(left, m, n)
			if !isOk{
				break
			}
		}
	}
	return false
}
