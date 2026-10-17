const Infinity = math.MaxInt
func minEatingSpeed(piles []int, h int) int {
	minVal, maxVal := 1, -Infinity
	for _, pile := range piles {
		maxVal = max(maxVal, pile)
	} 
	result := Infinity
	for minVal <= maxVal {
		midVal := (minVal + maxVal) / 2
		times := 0
		for _, pile := range piles{
			if pile % midVal == 0 {
				times += pile/midVal
			} else {
				times += pile/midVal + 1
			}
		}
		if times > h {
			minVal = midVal + 1
		} else  {
			result = min(result, midVal)
			maxVal = midVal - 1
		}
	}
	return result
}
