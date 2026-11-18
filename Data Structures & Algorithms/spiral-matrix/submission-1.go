import "math"
const Visited = math.MaxInt

type Direction int

const (
	Right Direction = iota
	Down
	Left
	Up
)

func (d Direction) Next() Direction {
	return (d + 1) % 4
}

func nextPosition(dir Direction, curPos []int, matrix [][]int) (Direction,[]int){
	curRow, curCol := curPos[0], curPos[1]
	nextRow, nextCol := curRow, curCol
	m,n := len(matrix),len(matrix[0])
	if dir == Up {
		nextRow = curRow - 1
	} else if dir == Down{
		nextRow = curRow + 1
	} else if dir == Left{
		nextCol = curCol - 1
	} else {
		nextCol = curCol + 1
	}
	if nextRow < 0 || nextRow >= m || nextCol < 0 || nextCol >= n || matrix[nextRow][nextCol] == Visited{
		dir = dir.Next()
	} else {
		return dir, []int{nextRow, nextCol}
	}
	nextRow, nextCol = curRow, curCol
	if dir == Up {
		nextRow = curRow - 1
	} else if dir == Down{
		nextRow = curRow + 1
	} else if dir == Left{
		nextCol = curCol - 1
	} else {
		nextCol = curCol + 1
	}
	return dir, []int{nextRow, nextCol}
}

func spiralOrder(matrix [][]int) []int {
    m, n := len(matrix),len(matrix[0])
	curPos := []int{0,0}
	dir := Right
	result := make([]int,0)
	for i := 0; i < m*n ; i++{
		row, col := curPos[0], curPos[1]
		result = append(result, matrix[row][col])
		matrix[row][col]=Visited
		dir, curPos = nextPosition(dir, curPos, matrix)
	}
	return result
}
