// dp[x]: number of way to x
// dp[x] = dp[x-1] + dp[x-2] + dp[x-3]
// dp[1][1] = 1
// dp[1][2] = 1
// dp[1][3] = 1
// dp[2][1] = 1 create i by first j element
// dp[2][2] = 2 => create i by first j-1 element + 1 if i - coins[j] == 0
// dp[2][3] = 2
// dp[3][1] = 0 + dp[2][1] = 1
// dp[3][2] = 1 + dp[1][2] = 2
// dp[3][3] = 2 + dp[0][3] = 3
// dp[4][1] = 0 + dp[3][1] = 1
// dp[4][2] = 1 + dp[2][2] = 3



func change(amount int, coins []int) int {
	if amount == 0 {return 1}
    n := len(coins)
	dp := make([][]int, (amount+1)*(n+1))
	for row := range dp{
		dp[row]=make([]int, n+1)
	}
	for i := 1; i <= amount; i++ {
		for j := 1; j <= n; j++ {
			tmp := dp[i][j-1]
			prev := i - coins[j-1]
			if prev == 0 {
				tmp = tmp + 1
			} else if prev > 0 {
				tmp = tmp + dp[prev][j]
			}
			dp[i][j]=tmp
		}
	}
	return dp[amount][n]
}
