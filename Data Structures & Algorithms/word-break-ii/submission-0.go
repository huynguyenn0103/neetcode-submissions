import "slices"

func wordBreak(s string, wordDict []string) []string {
	dict := make(map[string]bool)
	for _, word := range wordDict{
		dict[word]=true
	}
	result := make([][]string,0)
	combination := make([]string,0)
	n := len(s)
	var backtrack func(idx int)
	backtrack = func(idx int){
		if idx >= n {
			cloneCombination := slices.Clone(combination)
			result = append(result, cloneCombination)
			return
		}
		for i:=idx; i<=n; i++ {
			subStr := s[idx:i]
			if !dict[subStr]{continue}
			combination = append(combination,subStr)
			backtrack(i)
			combination = combination[:len(combination)-1]
		}
	}
	backtrack(0)
	fResult := make([]string,0)

	for _, strList := range result{
		var builder strings.Builder
		for i, str := range strList{
			builder.WriteString(str)
			if i != len(strList)-1{
				builder.WriteString(" ")
			}
		}
		fResult = append(fResult,builder.String())
	}
	return fResult
}
