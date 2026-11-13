func leastBricks(wall [][]int) int {
    mapBoundaryToFreq := make(map[int]int)
	for _, w := range wall{
		sum := 0
		n := len(w)
		for i := 0; i < n-1; i++{
			sum = sum + w[i]
			mapBoundaryToFreq[sum]++
		}
	}
	maxFreq := 0
	for _, freq := range mapBoundaryToFreq{
		maxFreq = max(maxFreq, freq)
	}
	return len(wall) - maxFreq
}