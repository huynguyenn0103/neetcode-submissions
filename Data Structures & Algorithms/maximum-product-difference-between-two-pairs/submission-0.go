func maxProductDifference(nums []int) int {
    firstMax, secondMax := -math.MaxInt, -math.MaxInt
	firstMin, secondMin := math.MaxInt, math.MaxInt
	for i := 0; i < len(nums); i++ {
		if nums[i] >= firstMax{
			secondMax = firstMax
			firstMax = nums[i]
		} else if nums[i] > secondMax{
			secondMax = nums[i]
		}
		if nums[i] <= firstMin{
			secondMin = firstMin
			firstMin = nums[i]
		} else if nums[i] < secondMin{
			secondMin = nums[i]
		}
	}
	return (firstMax * secondMax) - (firstMin * secondMin)
}