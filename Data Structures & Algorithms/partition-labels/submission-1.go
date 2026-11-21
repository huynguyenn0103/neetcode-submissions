func partitionLabels(s string) []int {
	mapByteToLastIdx := make(map[byte]int)
	n := len(s)
	for i := 0; i < n; i++{
		c := s[i]
		mapByteToLastIdx[c] = i
	}
	result := make([]int,0)
	maxIdx := -1
	left := 0
	for right := 0; right < n; right++{
		c := s[right]
		maxIdx = max(maxIdx, mapByteToLastIdx[c])
		if maxIdx == right{
			result = append(result, right - left + 1)
			maxIdx = -1
			left = right + 1
		}
	}
	return result
}
