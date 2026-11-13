func subarraysDivByK(nums []int, k int) int {
  result := 0
  n := len(nums)
  for i := 0; i < n; i++{
	sum := 0
	for j := i; j < n; j++{
		sum = sum + nums[j]
		if sum % k == 0{
			result++
		}
	}
  } 
  return result
}