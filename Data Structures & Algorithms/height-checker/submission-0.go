import "slices"
func heightChecker(heights []int) int {
	tmp := slices.Clone(heights)
    slices.SortFunc(heights, func(h1,h2 int)int{return h1-h2})
	result := 0
	for i := 0; i < len(tmp); i++ {
		if tmp[i] != heights[i]{
			result++
		}
	}
	return result
}