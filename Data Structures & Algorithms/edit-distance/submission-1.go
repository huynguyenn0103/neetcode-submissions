func minDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)

    memo := make(map[[2]int]int)

    var dfs func(i, j int) int

    dfs = func(i, j int) int {
        // word1 is exhausted: insert remaining characters
        if i == m {
            return n - j
        }
        // word2 is exhausted: delete remaining characters
        if j == n {
            return m - i
        }
        key := [2]int{i, j}
        if result, exists := memo[key]; exists {
            return result
        }
        if word1[i] == word2[j] {
            memo[key] = dfs(i+1, j+1)
        } else {
            insert := dfs(i, j+1)
            delete := dfs(i+1, j)
            replace := dfs(i+1, j+1)
            memo[key] = 1 + min(insert, min(delete, replace))
        }
        return memo[key]
    }
    return dfs(0, 0)
}