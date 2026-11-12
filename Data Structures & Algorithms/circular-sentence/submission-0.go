func isCircularSentence(sentence string) bool {
	strs := strings.Split(sentence, " ")
	if len(strs) == 1{
		str := strs[0]
		return str[0] == str[len(str)-1]
	}
	n := len(strs)
	if strs[0][0] != strs[n-1][len(strs[n-1])-1]{
		return false
	}
	for i := 1; i < n; i++ {
		last := strs[i-1]
		cur := strs[i]
		if cur[0] != last[len(last)-1]{return false}
	}
	return true
}
