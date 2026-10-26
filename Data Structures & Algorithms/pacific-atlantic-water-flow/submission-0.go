const VisitedPacific = 1
const Visited = 2

type Point struct {
	row int
	col int
}

var rowActions = []int{-1,0,1,0}
var colActions = []int{0,1,0,-1}

func pacificAtlantic(heights [][]int) [][]int {
    result := make([][]int, 0)
	queueP := make([]Point, 0)
	queueA := make([]Point, 0)
	m, n := len(heights), len(heights[0])
	visited := make([][]int, m*n)
	for row := range visited{
		visited[row]=make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				queueP = append(queueP, Point{row: i, col: j})
			}
			if i == m-1 || j == n-1 {
				queueA = append(queueA, Point{row: i, col: j})
			}
		}
	}
	for len(queueP) > 0 {
		point := queueP[0]
		queueP = queueP[1:]
		row, col := point.row, point.col
		if visited[row][col] == VisitedPacific{continue}
		visited[row][col]=VisitedPacific
		for i := 0; i < len(rowActions); i++ {
			nextRow := row + rowActions[i]
			nextCol := col + colActions[i]
			if nextRow >= 0 && nextRow < m && nextCol >= 0 && nextCol < n && heights[nextRow][nextCol] >= heights[row][col]{
				queueP = append(queueP, Point{row: nextRow, col: nextCol})
			}
		}
	}
	for len(queueA) > 0 {
		point := queueA[0]
		queueA = queueA[1:]
		row,col := point.row, point.col
		if visited[row][col] == Visited{continue}
		if visited[row][col] == VisitedPacific{
			result = append(result, []int{row,col})
		}
		visited[row][col]=Visited
		for i := 0; i < len(rowActions); i++ {
			nextRow := row + rowActions[i]
			nextCol := col + colActions[i]
			if nextRow >= 0 && nextRow < m && nextCol >= 0 && nextCol < n && heights[nextRow][nextCol] >= heights[row][col]{
				queueA = append(queueA, Point{row: nextRow, col: nextCol})
			}
		}
	}
	return result
}
