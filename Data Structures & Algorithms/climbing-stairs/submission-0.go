func climbStairs(n int) int {
    result := make([]int, n+1)
	for i := 0; i <= n; i++ {
		if i == 1 || i == 0{
			result[i] = 1
		} else {
			result[i] = result[i-1] + result[i-2]
		}
	}
	return result[n]
}
