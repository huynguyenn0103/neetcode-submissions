func getRow(rowIndex int) []int {
    curRow := []int{1}
	for rowIndex > 0{
		tmp := make([]int,0)
		tmp = append(tmp, 1)
		for i := 1; i < len(curRow); i++ {
			tmp = append(tmp, curRow[i]+curRow[i-1])
		}
		tmp = append(tmp, 1)
		curRow = tmp
		rowIndex--
	}
	return curRow
}