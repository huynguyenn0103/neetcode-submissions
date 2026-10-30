// aabb  bbcc
// aa bbcc bb

//dp[0][1]=false
//dp[1][0]=true

//dp[0][2]=false based on dp[0][1] & b==a?
//dp[1][1]=false based on dp[0][1] && a==a? || dp[1][0] && b==a?
//dp[2][0]=true based on dp[1][0] && a==a

//dp[0][3]
//dp[1][2] based on dp[1][1] & b==b || dp[0][2] && a==a



func isInterleave(s1 string, s2 string, s3 string) bool {
    m,n,l := len(s1), len(s2), len(s3)
	if m + n != l {return false}
	if l == 0 {return true}
	dp := make([][]bool, (m+1)*(n+1))
	for row := range dp{
		dp[row]=make([]bool, n+1)
	}
	dp[0][0]=true
	for k:=1; k<=l; k++ {
		for i:=0; i<=k && i<=m; i++ {
			j:=k-i
			if j > n {continue}
			isSuccess:= false
			if i >= 1 && s1[i-1]==s3[k-1] && dp[i-1][j]{
				isSuccess = true
			}
			if j >= 1 && s2[j-1]==s3[k-1] && dp[i][j-1]{
				isSuccess = true
			}
			dp[i][j]=isSuccess
		}
	}
	return dp[m][n]
}
