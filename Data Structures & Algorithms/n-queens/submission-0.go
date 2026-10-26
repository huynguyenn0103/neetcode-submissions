import "slices"

type Point struct{
	row int
	col int
}

func dist(v1, v2 int)int{
	return int(math.Abs(float64(v1)-float64(v2)))
}

func isViolatedTwoQueen(p1, p2 Point)bool {
	x1,y1,x2,y2 := p1.row, p1.col, p2.row, p2.col
	if x1 == x2 || y1 == y2 || dist(x1,x2) == dist(y1,y2) {
		return true
	}
	return false
}

func isViolated(p Point, curQueens []Point)bool {
	for _, queen := range curQueens {
		if isViolatedTwoQueen(p, queen){
			return true
		}
	}
	return false
}

func convertPointsToResult(points [][]Point) [][]string{
	result := make([][]string, 0)
	if len(points) == 0 {return result}
	n := len(points[0])
	for _, point := range points {
		tmp := make([]string, 0)
		for i := 0; i < n; i++ {
			var builder strings.Builder
			for j := 0; j < n; j++ {
				if len(point) > 0 {
					top := point[0]
					if top.row == i && top.col == j{
						builder.WriteString("Q")
						point = point[1:]
					} else {
						builder.WriteString(".")
					}
				} else {
					builder.WriteString(".")
				}
			}
			tmp = append(tmp, builder.String())
		}
		result = append(result, tmp)
	}
	return result
}

func solveNQueens(n int) [][]string {
	resultInPoints := make([][]Point, 0)
	combination := make([]Point, 0)
	puzzle := make([][]int, n * n)
	for row := range puzzle {
		puzzle[row] = make([]int, n)
	}
	var backtrack func(row, col int)
	backtrack = func(row, col int){
		if len(combination) == n {
			resultInPoints = append(resultInPoints, slices.Clone(combination))
			return
		}
		if row >= n || col >= n {return}
		for row < n {
			for col < n {
				candidate := Point{row,col}
				if !isViolated(candidate, combination){
					combination = append(combination, candidate)
					var nextRow, nextCol int
					if col + 1 >= n {
						nextRow = row + 1
						nextCol = 0
					} else {
						nextRow = row
						nextCol = col + 1
					}
					backtrack(nextRow, nextCol)
					combination = combination[:len(combination)-1]
				}
				col++
			}
			col = 0
			row++
		}
	}
	backtrack(0,0)
	return convertPointsToResult(resultInPoints)
}
