func kthDistinct(arr []string, k int) string {
    mapStr := make(map[string]int)
	for _, str := range arr{
		mapStr[str]++
	}
	for _, str := range arr {
		if mapStr[str] == 1 {
			k--
		}
		if k == 0 {
			return str
		}
	}
	return ""
}