func characterReplacement(s string, k int) int {
	result := 0
	for target := byte('A'); target <= byte('Z'); target++{
		updatedIdx := make([]int, 0)
		start := 0
		n := len(s)
		for i := 0; i < n; i++ {
			c := s[i]
			if c == target {continue}
			if c != target{
				if len(updatedIdx) < k {
					updatedIdx = append(updatedIdx, i)
				} else {
					result = max(result, i - start)
					if k > 0 {
						removeIdx := updatedIdx[0]
						updatedIdx = updatedIdx[1:]
						updatedIdx = append(updatedIdx,i)
						start = removeIdx + 1
					} else {
						start = i + 1
					}
				}
			}
		}
		result = max(result, n - start)
	}
	return result
}
