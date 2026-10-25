func generateParenthesis(n int) []string {
    result := make([]string, 0)
    combination := make([]byte, 0, 2*n)

    var backtrack func(open, close int)

    backtrack = func(open, close int) {
        if len(combination) == 2*n {
            result = append(result, string(combination))
            return
        }

        // Add '(' if we still have opening parentheses available
        if open < n {
            combination = append(combination, '(')
            backtrack(open+1, close)
            combination = combination[:len(combination)-1]
        }

        // Add ')' only if it doesn't make the combination invalid
        if close < open {
            combination = append(combination, ')')
            backtrack(open, close+1)
            combination = combination[:len(combination)-1]
        }
    }

    backtrack(0, 0)
    return result
}