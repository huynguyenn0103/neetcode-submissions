func minCostClimbingStairs(cost []int) int {
    n := len(cost)
	result := make([]int,n+1)
	for i := 2; i <= n; i++ {
		result[i]=min(result[i-1]+cost[i-1], result[i-2]+cost[i-2])
	}
	return result[n]
}
