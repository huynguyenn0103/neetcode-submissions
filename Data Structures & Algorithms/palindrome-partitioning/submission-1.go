// aabb

type CacheResult struct{
	result [][]string
	isSuccess bool
}

func isPalindrome(str string)bool{
	n := len(str)
	for i := 0; i < n/2; i++{
		if str[i] != str[n-1-i]{
			return false
		}
	}
	return true
}

func partition(s string) [][]string {
	cacheResultFromIdx := make(map[int]CacheResult)
	n := len(s)

	var buildPalindromeSubstrFromIdx func(idx int)CacheResult
	buildPalindromeSubstrFromIdx = func(idx int)CacheResult{
		if idx >= n {
			return CacheResult{result: [][]string{}, isSuccess: true}
		}
		if result, existed := cacheResultFromIdx[idx]; existed{
			return result
		}
		result := make([][]string,0)
		isSuccess := false
		curBytes := make([]byte,0)
		for i := idx; i < n; i++ {
			curBytes = append(curBytes, s[i])
			if !isPalindrome(string(curBytes)){continue}
			palindromesFromNextResult := buildPalindromeSubstrFromIdx(i+1)
			if !palindromesFromNextResult.isSuccess{continue}
			isSuccess = true
			palindromes := palindromesFromNextResult.result
			if len(palindromes) == 0 {
				result = append(result, []string{string(curBytes)})
			} else {
				for _, palindrome := range palindromes{
					tmp := make([]string,0)
					tmp = append(tmp, string(curBytes))
					tmp = append(tmp, palindrome...)
					result = append(result, tmp)
				}
			}
		}
		cacheResultFromIdx[idx] = CacheResult{
			result: result,
			isSuccess: isSuccess,
		}
		return cacheResultFromIdx[idx]
	}
	result := buildPalindromeSubstrFromIdx(0)
	return result.result
}
