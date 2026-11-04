func simplifyPath(path string) string {
	components := strings.Split(path,"/")
	stack := make([]string, 0)
	for _, component := range components {
		if len(component) == 0 || component == "."{continue}
		if component == ".."{
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack,component)
		}
	}
	if len(stack) == 0 {return "/"}
	var builder strings.Builder
	for _, component := range stack {
		builder.WriteString("/")
		builder.WriteString(component)
	}
	return builder.String()
}
