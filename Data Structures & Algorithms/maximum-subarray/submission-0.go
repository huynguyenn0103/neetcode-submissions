func maxSubArray(nums []int) int {
   result := nums[0]
   sum := nums[0]
   for i := 1; i < len(nums); i++ {
	 sum = sum + nums[i]
	 sum = max(sum, nums[i])
	 result = max(sum, result)
   }
   return result
}
