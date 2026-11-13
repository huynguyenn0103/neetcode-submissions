func appendCharacters(s string, t string) int {
    ptr1, ptr2 := 0, 0
	m, n := len(s), len(t)
	for ptr1 < m && ptr2 < n {
		if t[ptr2] == s[ptr1]{
			ptr2++
		}
		ptr1++
	}
	return len(t) - ptr2
}