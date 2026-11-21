import "slices"

func generateParenthesis(n int) []string {
	curOpen := 0
	curClose := 0
	combination := make([]byte,0)
	result := make([]string,0)

	var buildParenthesis func()
	buildParenthesis = func(){
		if len(combination) == 2*n{
			clone := slices.Clone(combination)
			result = append(result, string(clone))
			return
		}
		if curOpen == n || curOpen > curClose{
			combination = append(combination, ')')
			curClose++
			buildParenthesis()
			curClose--
			combination = combination[:len(combination)-1]
		}
		
		if curOpen == curClose || (curOpen > curClose && curOpen != n){
			combination = append(combination, '(')
			curOpen++
			buildParenthesis()
			curOpen--
			combination = combination[:len(combination)-1]
		}
	}
	buildParenthesis()
	return result
}
