func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)

    // Initialize all amounts as impossible
    for i := 1; i <= amount; i++ {
        dp[i] = amount + 1
    }

    // Base case
    dp[0] = 0

    for currAmount := 1; currAmount <= amount; currAmount++ {
        for _, coin := range coins {
            if coin <= currAmount {
                dp[currAmount] = min(
                    dp[currAmount],
                    dp[currAmount-coin]+1,
                )
            }
        }
    }

    if dp[amount] == amount+1 {
        return -1
    }

    return dp[amount]
}