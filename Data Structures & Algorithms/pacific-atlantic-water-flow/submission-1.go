const Pacific = 1
const Atlantic = 2
const Unvisited = 0

var rowActions = []int{-1,0,1,0}
var colActions = []int{0,1,0,-1}

type Cell struct{
	row int
	col int
}

func pacificAtlantic(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])
	water := make([][]int,m)
	for row := range water{
		water[row]=make([]int,n)
	}
	for i := 0; i < m; i++{
		for j := 0; j < n; j++{
			if (i == 0 || j == 0) && water[i][j] == Unvisited{
				queue := make([]Cell,0)
				queue = append(queue, Cell{row: i, col: j})
				for len(queue) > 0 {
					cell := queue[0]
					queue = queue[1:]
					row,col := cell.row, cell.col
					if water[row][col] != Unvisited{continue}
					water[row][col]=Pacific
					for i := 0; i < len(rowActions); i++{
						nextRow := row + rowActions[i]
						nextCol := col + colActions[i]
						if nextRow >= 0 && nextRow < m && nextCol >= 0 && nextCol < n && heights[row][col] <= heights[nextRow][nextCol]{
							queue = append(queue, Cell{row: nextRow, col: nextCol})
						}
					}
				}
			}
		}
	}

	result := make([][]int,0)
	for i := 0; i < m; i++{
		for j := 0; j < n; j++ {
			if (i == m-1 || j == n-1) && water[i][j] != Atlantic {
				queue := make([]Cell,0)
				queue = append(queue, Cell{row: i, col: j})
				for len(queue)>0{
					cell := queue[0]
					queue = queue[1:]
					row,col := cell.row,cell.col
					if water[row][col] == Atlantic{continue}
					if water[row][col] == Pacific{
						result = append(result, []int{row,col})
					}
					water[row][col] = Atlantic
					for i := 0; i < len(rowActions); i++{
						nextRow := row + rowActions[i]
						nextCol := col + colActions[i]
						if nextRow >= 0 && nextRow < m && nextCol >= 0 && nextCol < n && heights[row][col] <= heights[nextRow][nextCol]{
							queue = append(queue, Cell{row: nextRow, col: nextCol})
						}
					}
				}
			}
		}
	}
	return result
}
