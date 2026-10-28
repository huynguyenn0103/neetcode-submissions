// 2 9 8 100 6 1 3 5 6

// 2 -> 8/100 -> 6,1/1,5

func rob(nums []int) int {
   n := len(nums)
   maxRobTilIdx := make([]int,n) 
   for i := 0; i < n; i++ {
	   if i <= 1 {
			maxRobTilIdx[i]=nums[i]
	   } else if i == 2 {
			maxRobTilIdx[i]=nums[i]+nums[i-2]
	   } else {
			maxRobTilIdx[i]=max(maxRobTilIdx[i-2],maxRobTilIdx[i-3])+nums[i]
	   }
   }
   if n == 1{
	return maxRobTilIdx[n-1]
   }
   return max(maxRobTilIdx[n-2],maxRobTilIdx[n-1])
}
