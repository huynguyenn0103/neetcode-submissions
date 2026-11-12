func minOperations(s string) int {
	operationWithStartZero := 0
	operationWithStartOne := 0
	cur := byte('0')
	for i := 0; i < len(s); i++{
		if s[i] == cur{
			operationWithStartOne++
		} else {
			operationWithStartZero++
		}
		cur = '0' + ('1' - cur)
	}
	return min(operationWithStartZero, operationWithStartOne)
}