import "slices"
func relativeSortArray(arr1 []int, arr2 []int) []int {
    mapArr1ToFreq := make(map[int]int)
	for _, arr := range arr1{
		mapArr1ToFreq[arr]++
	}
	result := make([]int, 0)
	for _, arr := range arr2 {
		for mapArr1ToFreq[arr] > 0 {
			result = append(result, arr)
			mapArr1ToFreq[arr]--
		}
	}
	remain := make([]int, 0)
	for key := range mapArr1ToFreq{
		for mapArr1ToFreq[key] > 0 {
			remain = append(remain, key)
			mapArr1ToFreq[key]--
		}
	}
	slices.SortFunc(remain, func(i1, i2 int)int{return i1 - i2})
	result = append(result, remain...)
	return result
}