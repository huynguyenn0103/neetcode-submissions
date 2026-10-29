
// dp[idx]=true means can split til idx is end
func wordBreak(s string, wordDict []string) bool {
   dict := make(map[string]bool)
   for _, word := range wordDict{
	dict[word]=true
   } 
   n := len(s)
   dp := make([]bool, n)
   for i := 0; i < n; i++{
	 isSuccess := false
	 for j := i; j >= 0; j--{
		if dict[s[j:i+1]] && (j == 0 || (j > 0 && dp[j-1])){
			isSuccess = true
			break
		}
	 }
	 dp[i]=isSuccess
   }
   return dp[n-1]
}
