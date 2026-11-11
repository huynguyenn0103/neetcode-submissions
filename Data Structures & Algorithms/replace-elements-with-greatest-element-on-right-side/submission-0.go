func replaceElements(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		maxVal := arr[i+1]
		for j := i + 1; j < n; j++ {
			maxVal = max(maxVal, arr[j])
		}
		arr[i] = maxVal
	}
	arr[n-1] = -1
	return arr
}
