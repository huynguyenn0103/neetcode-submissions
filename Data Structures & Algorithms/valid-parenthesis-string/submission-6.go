// ())

func checkValidString(s string) bool {
    numOfStackIfAllStarIsOpen := 0
	numOfStackIfAllStarIsCloseOrEmpty := 0
	n := len(s)
	for i := 0; i < n; i++ {
		c := s[i]
		switch c {
			case '(':
				numOfStackIfAllStarIsOpen++
				numOfStackIfAllStarIsCloseOrEmpty++
			case ')':
				numOfStackIfAllStarIsOpen--
				numOfStackIfAllStarIsCloseOrEmpty--
			case '*':
				numOfStackIfAllStarIsCloseOrEmpty--
				numOfStackIfAllStarIsOpen++
		}
		if numOfStackIfAllStarIsOpen < 0 {
			return false
		}
		if numOfStackIfAllStarIsCloseOrEmpty < 0 {
			numOfStackIfAllStarIsCloseOrEmpty = 0
		}
	}
	return numOfStackIfAllStarIsCloseOrEmpty == 0
}
