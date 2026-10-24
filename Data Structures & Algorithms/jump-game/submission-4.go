// 5 9 3 2 1 0 2 3 3 1 0 0
func canJump(nums []int) bool {
   n := len(nums)
   if n == 1{
	 return true
   }
   start, end := 0, nums[0]
   for {
	  next := 0
	  for start <= end {
		next = max(next, start + nums[start])
		if next >= n-1{
			return true
		}
		start++
	  }
	  if next <= end {
		return false
	  }
	  end = next
   }
   return false
}

//1 5 7 0 6 0 0 0 0 0 0
