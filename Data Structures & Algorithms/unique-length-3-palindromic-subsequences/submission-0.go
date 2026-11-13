func numOfDistinctVal(left, right int, s string)int{
	visited := make(map[byte]bool)
	for i := left + 1; i < right; i++{
		visited[s[i]]=true
	}
	return len(visited)
}

func countPalindromicSubsequence(s string) int {
	mapByteToFirstIdx := make(map[byte]int)
	mapByteToLastIdx := make(map[byte]int)
	n := len(s)
	for i := 0; i < n; i++{
		c := s[i]
		if _, existed := mapByteToFirstIdx[c]; !existed{
			mapByteToFirstIdx[c] = i
		}
		mapByteToLastIdx[c] = i
	}

	result := 0

	for c, left := range mapByteToFirstIdx{
		result += numOfDistinctVal(left, mapByteToLastIdx[c],s)
	}
	return result
}