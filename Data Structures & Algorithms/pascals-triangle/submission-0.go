func generate(numRows int) [][]int {
	result := make([][]int,0)
	result = append(result, []int{1})
	for numRows > 1{
		lastArr := result[len(result)-1]
		n := len(lastArr)
		tmp := make([]int, 0)
		tmp = append(tmp, 1)
		for i := 0; i < n - 1; i++ {
			tmp = append(tmp, lastArr[i] + lastArr[i+1])
		}
		tmp = append(tmp, 1)
		result = append(result, tmp)
		numRows--
	}
	return result
}
