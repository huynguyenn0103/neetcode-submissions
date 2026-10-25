import "slices"
func isPalindrome(str string)bool{
	n := len(str)
	for i := 0; i < n/2; i++{
		if str[i] != str[n-1-i]{
			return false
		}
	}
	return true
}
func partition(s string) [][]string {
	result := make([][]string,0)
	combination := make([]string, 0)
	var backtrack func(fromIdx int)
	backtrack = func(fromIdx int){
		if fromIdx == len(s){
			result = append(result, slices.Clone(combination))
			return
		}
		var builder strings.Builder
		for i := fromIdx; i < len(s); i++ {
			builder.WriteString(string(s[i]))
			if isPalindrome(builder.String()){
				combination = append(combination, builder.String())
				backtrack(i+1)
				combination = combination[:len(combination)-1]
			} 
		}
	}
	backtrack(0)
	return result
}
