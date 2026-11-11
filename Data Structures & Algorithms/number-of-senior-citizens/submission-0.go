func countSeniors(details []string) int {
    result := 0 
	for _, detail := range details {
		num, _ := strconv.Atoi(detail[11:13])
		if num > 60 {
			result++
		}
	}
	return result
}