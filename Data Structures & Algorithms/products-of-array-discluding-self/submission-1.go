//  1 1 2 8
// 48 24 6 1

func productExceptSelf(nums []int) []int {
	n := len(nums)
	multiplicationBeforeIdx := make([]int,0)
	multiplicationBeforeIdx = append(multiplicationBeforeIdx,1)
	mul := 1
	for i := 0; i < n-1; i++{
		mul = mul * nums[i]
		multiplicationBeforeIdx = append(multiplicationBeforeIdx,mul)
	}

	mul = 1
	multiplicationAfterIdx := make([]int,0)
	multiplicationAfterIdx = append(multiplicationAfterIdx,1)
	for i := n-1; i > 0; i--{
		mul = mul * nums[i]
		multiplicationAfterIdx = append(multiplicationAfterIdx,mul)
	}
	result := make([]int,0)
	for i := 0; i < n; i++ {
		result = append(result, multiplicationBeforeIdx[i] * multiplicationAfterIdx[n-1-i])
	}
	return result
}
