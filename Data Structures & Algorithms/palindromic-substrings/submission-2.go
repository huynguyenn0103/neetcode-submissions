func countSubstrings(s string) int {
	n := len(s)
    dp := make([][]bool,n)
	result := 0
	for idx := range dp{
		dp[idx]=make([]bool,n)
	}
	for i := n-1; i >= 0; i--{
		for j := i; j < n; j++{
			if s[i]==s[j] && (j-i<=2 || dp[i+1][j-1]){
				dp[i][j]=true
				result++
			}
		}
	}
	return result
}
