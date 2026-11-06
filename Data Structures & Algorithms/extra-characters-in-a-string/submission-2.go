func minExtraChar(s string, dictionary []string) int {
    n := len(s)

    wordSet := make(map[string]bool)
    for _, word := range dictionary {
        wordSet[word] = true
    }

    memo := make([]int, n+1)
    for i := range memo {
        memo[i] = -1
    }

    var dfs func(i int) int
    dfs = func(i int) int {
        if i == n {
            return 0
        }

        if memo[i] != -1 {
            return memo[i]
        }

        // Option 1: Skip current character
        result := 1 + dfs(i+1)

        // Option 2: Match a dictionary word starting at i
        for j := i + 1; j <= n; j++ {
            word := s[i:j]

            if wordSet[word] {
                result = min(result, dfs(j))
            }
        }

        memo[i] = result
        return result
    }

    return dfs(0)
}