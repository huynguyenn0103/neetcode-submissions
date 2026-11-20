import "slices"

func isOverlap(i1, i2 []int)bool{
	return i1[1] >= i2[0]
}

func mergeInterval(i1, i2 []int)[]int{
	return []int{i1[0], max(i1[1],i2[1])}
}


func merge(intervals [][]int) [][]int {
    slices.SortFunc(intervals, func(i1,i2 []int)int{
		if i1[0] == i2[0]{
			return i1[1] - i2[1]
		}
		return i1[0] - i2[0]
	})
	result := make([][]int,0)
	for _, interval := range intervals{
		if len(result) == 0{
			result = append(result, interval)
		} else {
			lastInterval := result[len(result)-1]
			if isOverlap(lastInterval, interval){
				m := mergeInterval(lastInterval, interval)
				result = result[:len(result)-1]
				result = append(result, m)
			} else {
				result = append(result, interval)
			}
		}
	}
	return result
}
