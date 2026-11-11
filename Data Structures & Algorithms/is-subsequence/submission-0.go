func isSubsequence(s string, t string) bool {
	ptr1, ptr2 := 0,0
	for ptr1 < len(s) && ptr2 < len(t){
		if s[ptr1] == t[ptr2]{
			ptr1++
		}
		ptr2++
	}
	return ptr1 == len(s)
}
