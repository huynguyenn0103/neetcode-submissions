func maxCoins(nums []int) int {
    n := len(nums)

    arr := make([]int, n+2)
    arr[0], arr[n+1] = 1, 1
    copy(arr[1:], nums)

    dp := make([][]int, n+2)
    for i := range dp {
        dp[i] = make([]int, n+2)
        for j := range dp[i] {
            dp[i][j] = -1
        }
    }

    var dfs func(left, right int) int

    dfs = func(left, right int) int {
        if right-left <= 1 {
            return 0
        }

        if dp[left][right] != -1 {
            return dp[left][right]
        }

        maxCoins := 0

        for k := left + 1; k < right; k++ {
            coins := arr[left] * arr[k] * arr[right]

            coins += dfs(left, k)
            coins += dfs(k, right)

            maxCoins = max(maxCoins, coins)
        }

        dp[left][right] = maxCoins
        return maxCoins
    }

    return dfs(0, n+1)
}