const Inf = math.MaxInt

func minSubArrayLen(target int, nums []int) int {
	if nums[0] >= target {return 1}
	if len(nums) == 1 {
		return 0
	}
	n := len(nums)
	result := Inf
	left, right := 0, 1
	sum := nums[left]
	
	for right < n {
		sum += nums[right]
		if sum >= target {
			for {
				sum = sum - nums[left]
				if sum < target {
					result = min(result, right-left+1)
					left++
					if left < n {
						sum = nums[left]
					}
					right = left + 1
					break
				}
				left++
			}
		} else {
			right++
		}
	}
	if result == Inf {return 0}
	return result
}
