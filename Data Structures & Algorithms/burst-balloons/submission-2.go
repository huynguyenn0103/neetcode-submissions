// 1 a b c d 1

//2 3 4
// 1*2*3 -> 3 4
// 1*3*4 -> 4
// 1*4*1 = 22

// 2*3*4 -> 2 4
// 1*2*4 ->4
// 1*4*1

func maxCoins(nums []int) int {
    cache := make(map[[2]int]int)
	n := len(nums)
	var findMaxCoinFrom func(left, right int)int
	findMaxCoinFrom = func(left, right int)int{
		if left > right {
			return 0
		}
		key := [2]int{left,right}
		if val, existed := cache[key]; existed{
			return val
		}
		result := 0
		prev, after := 1,1
		if left >= 1{
			prev = nums[left-1]
		}
		if right < n-1{
			after = nums[right+1]
		}
		for i := left; i <= right; i++{
			result = max(result, nums[i]*prev*after + findMaxCoinFrom(left,i-1) + findMaxCoinFrom(i+1,right))
		}
		cache[key]=result
		return result
	}
	return findMaxCoinFrom(0,n-1)
}
