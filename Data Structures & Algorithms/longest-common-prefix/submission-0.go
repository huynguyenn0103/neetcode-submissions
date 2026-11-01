func getPrefix(str1, str2 string) string{
	var builder strings.Builder
	for i := 0; i < len(str1) && i < len(str2); i++ {
		if str1[i] != str2[i]{break}
		builder.WriteString(string(str1[i]))
	}
	return builder.String()
}

func longestCommonPrefix(strs []string) string {
    if len(strs) == 1 {return strs[0]}
	result := getPrefix(strs[0], strs[1])
	for i := 2; i < len(strs); i++ {
		result = getPrefix(result, strs[i])
	}
	return result
}
