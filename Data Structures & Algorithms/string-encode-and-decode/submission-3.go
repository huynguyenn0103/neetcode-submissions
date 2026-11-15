type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var builder strings.Builder
	for _, str := range strs{
		m := len(str)
		for j := 0; j < m; j++{
			c := str[j]
			toStr := strconv.Itoa(int(c))
			builder.WriteString(toStr)
			if j != m-1{
				builder.WriteString(",")
			}
		}
		builder.WriteString("|")
	}
	return builder.String()
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == ""{
		return []string{}
	}
	result := make([]string,0)
	strs := strings.Split(encoded, "|")
	strs = strs[:len(strs)-1]
	for _, str := range strs{
		var builder strings.Builder
		nums := strings.Split(str, ",")
		for _, num := range nums{
			numInInt, err := strconv.Atoi(num)
			if err == nil {
				builder.WriteString(string(numInInt))
			}
		}
		result = append(result, builder.String())
	}
	return result
}
