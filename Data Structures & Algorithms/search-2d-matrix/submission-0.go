func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1
	for left <= right{
		mid := (left + right)/2
		x, y := mid/n, mid % n
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
