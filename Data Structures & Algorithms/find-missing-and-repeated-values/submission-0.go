func findMissingAndRepeatedValues(grid [][]int) []int {
    var a, b int
	mapNum := make(map[int]bool)
	n := len(grid)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if mapNum[grid[i][j]] == true{
				a = grid[i][j]
			} else {
				mapNum[grid[i][j]] = true
			}
		}
	}
	for i := 1; i <= n*n; i++ {
		if !mapNum[i]{
			b = i
			break
		}
	}
	return []int{a,b}
}