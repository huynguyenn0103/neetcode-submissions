func processLocal(localName string)string {
	bytes := []byte(localName)
	result := make([]byte, 0)
	for _, b := range bytes{
		if b == '+'{break}
		if b == '.'{continue}
		result = append(result, b)
	}
	return string(result)
}

func numUniqueEmails(emails []string) int {
    mapEmails := make(map[string]bool)
	for _, email := range emails{
		strs := strings.Split(email, "@")
		finEmail := processLocal(strs[0]) + "@" + strs[1]
		mapEmails[finEmail]=true
	}
	return len(mapEmails)
}