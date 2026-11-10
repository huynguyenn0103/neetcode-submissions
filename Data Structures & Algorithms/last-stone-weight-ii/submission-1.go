
func lastStoneWeightII(stones []int) int {
    total := 0
    for _, stone := range stones {
        total += stone
    }

    target := total / 2
    n := len(stones)

    // -1: not calculated
    memo := make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, target+1)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }

    var dfs func(i, remaining int) int

    dfs = func(i, remaining int) int {
        // Base cases
        if i == n || remaining == 0 {
            return 0
        }

        // Check memoization
        if memo[i][remaining] != -1 {
            return memo[i][remaining]
        }

        // Choice 1: Skip current stone
        skip := dfs(i+1, remaining)

        // Choice 2: Take current stone
        take := 0
        if stones[i] <= remaining {
            take = stones[i] + dfs(i+1, remaining-stones[i])
        }

        // Find maximum achievable subset sum
        memo[i][remaining] = max(skip, take)

        return memo[i][remaining]
    }

    best := dfs(0, target)

    return total - 2*best
}
