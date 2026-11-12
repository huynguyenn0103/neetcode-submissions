func isOdd(x int)bool{
	return x % 2 == 1
}
func isEven(x int)bool{
	return x % 2 == 0
}

func isArraySpecial(nums []int) bool {
    n := len(nums)
	for i := 1; i < n; i++ {
		if (isEven(nums[i-1]) && isEven(nums[i])) || (isOdd(nums[i-1]) && isOdd(nums[i])){
			return false
		}  
	}
	return true
}