// dpR[idx] => longest end at idx

func lengthOfLIS(nums []int) int {
	n := len(nums)
    dp := make([]int,n+1)
	for i := 0; i < n; i++ {
		count := 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j]{
				count = max(dp[j]+1,count)
			}
		}
		dp[i]=count
	}
	result := 0
	for _, val := range dp{
		result = max(result,val)
	}
	return result
}
