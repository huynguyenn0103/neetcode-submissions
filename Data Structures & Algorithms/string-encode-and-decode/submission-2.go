

type Solution struct{}

const SplitRune = ":"
const SplitString = "|"
const EmptyCase = "e"

func (s *Solution) Encode(strs []string) string {
	n := len(strs)
	if n == 0{
		return EmptyCase
	}
	var builder strings.Builder
	for i, str := range strs {
		for j, r := range str{
			builder.WriteString(strconv.Itoa(int(r)))
			if j != len(str) - 1 {
				builder.WriteString(SplitRune)
			}
		}
		if i != n-1 {
			builder.WriteString(SplitString)
		}
	}
	return builder.String()
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == EmptyCase{
		return []string{}
	}
	strs := strings.Split(encoded, SplitString)
	result := make([]string, 0)
	for _, str := range strs{
		runes := strings.Split(str, SplitRune)
		var builder strings.Builder
		for _, r := range runes{
			asInt, err := strconv.Atoi(r)
			if err == nil {
				builder.WriteString(string(asInt))
			}
		}
		result = append(result, builder.String())
	}
	return result
}
