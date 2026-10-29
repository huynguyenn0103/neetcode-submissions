func canPartition(nums []int) bool {
    total := 0

    for _, num := range nums {
        total += num
    }

    if total%2 != 0 {
        return false
    }

    target := total / 2
    n := len(nums)

    dp := make([][]bool, n+1)

    for i := range dp {
        dp[i] = make([]bool, target+1)
        dp[i][0] = true
    }

    for i := 1; i <= n; i++ {
        num := nums[i-1]

        for j := 1; j <= target; j++ {
            // Option 1: Skip current number
            dp[i][j] = dp[i-1][j]

            // Option 2: Take current number
            if j >= num {
                dp[i][j] = dp[i][j] || dp[i-1][j-num]
            }
        }
    }

    return dp[n][target]
}