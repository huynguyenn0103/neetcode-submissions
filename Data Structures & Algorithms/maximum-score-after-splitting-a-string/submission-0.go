func maxScore(s string) int {
	n := len(s)
    scoreLeft := 0
	if s[0] == '0'{
		scoreLeft++
	}

	scoreRight := 0
	for i := 1; i < n; i++ {
		if s[i] == '1'{
			scoreRight++
		}
	}
	result := scoreRight + scoreLeft
	for i := 1; i < n-1; i++ {
		if s[i] == '0' {
			scoreLeft++
		} else {
			scoreRight--
		}
		result = max(result, scoreLeft + scoreRight)
	}
	return result
}