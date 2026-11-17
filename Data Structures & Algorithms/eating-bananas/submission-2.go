func timeToEat(rate int, piles []int)int{
	result := 0
	for _, pile := range piles {
		if pile % rate == 0 {
			result = result + pile/rate
		} else {
			result = result + pile/rate + 1
		}
	}
	return result
}

func minEatingSpeed(piles []int, h int) int {
	maxPile := 0
	for _, pile := range piles{
		maxPile = max(maxPile, pile)
	}
	left, right := 1, maxPile
	result := math.MaxInt
	for left <= right{
		mid := (left + right)/2
		time := timeToEat(mid, piles)
		if time <= h{
			right = mid - 1
			result = min(result, mid)
		} else {
			left = mid + 1
		}
	}
	return result
}
