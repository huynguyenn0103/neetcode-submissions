func countConsistentStrings(allowed string, words []string) int {
    mapAllowed := make(map[byte]bool)
	for i := 0; i < len(allowed); i++ {
		mapAllowed[allowed[i]]=true
	}
	result := 0
	for _, word := range words{
		n := len(word)
		isConsistent := true
		for i := 0; i < n; i++ {
			if !mapAllowed[word[i]]{
				isConsistent = false
			}
		}
		if isConsistent{
			result++
		}
	}
	return result
}