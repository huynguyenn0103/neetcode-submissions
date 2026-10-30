func findTargetSumWays(nums []int, target int) int {
    dp := map[int]int{
        0: 1,
    }

    for _, num := range nums {
        next := make(map[int]int)

        for sum, count := range dp {
            next[sum+num] += count
            next[sum-num] += count
        }

        dp = next
    }

    return dp[target]
}