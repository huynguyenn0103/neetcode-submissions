func longestPalindrome(s string) string {
    n := len(s)

    dp := make([][]bool, n)
    for i := range dp {
        dp[i] = make([]bool, n)
    }

    start, maxLen := 0, 0

    for i := n - 1; i >= 0; i-- {
        for j := i; j < n; j++ {
            if s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1]) {
                dp[i][j] = true

                if j-i+1 > maxLen {
                    maxLen = j - i + 1
                    start = i
                }
            }
        }
    }

    return s[start : start+maxLen]
}