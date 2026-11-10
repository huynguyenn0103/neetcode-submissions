type Result struct {
	isLeft bool
	val int
}

func stoneGame(piles []int) bool {
	cacheMaxStone := make(map[[2]int]Result)
	var maxStoneFrom func(left, right int)Result
	maxStoneFrom = func(left, right int)Result{
		if left > right {
			return Result{isLeft: true, val: 0}
		}
		if left == right {
			return Result{isLeft: true, val: piles[left]}
		}
		if val, existed := cacheMaxStone[[2]int{left,right}]; existed{
			return val
		}

		result := 0

		bestOpponentIfChooseLeft := maxStoneFrom(left+1,right)
		bestOpponentIfChooseRight := maxStoneFrom(left, right-1)

		shouldLeft := true
		if bestOpponentIfChooseRight.val < bestOpponentIfChooseLeft.val{
			shouldLeft = false
		}
		if shouldLeft{
			if bestOpponentIfChooseLeft.isLeft{
				result = piles[left] + maxStoneFrom(left+2, right).val
			} else {
				result = piles[left] + maxStoneFrom(left+1, right-1).val
			}
		} else {
			if bestOpponentIfChooseRight.isLeft{
				result = piles[right] + maxStoneFrom(left+1, right-1).val
			} else {
				result = piles[right] + maxStoneFrom(left, right-2).val
			}
		}

		cacheMaxStone[[2]int{left,right}] = Result{isLeft: shouldLeft, val: result}
		return cacheMaxStone[[2]int{left,right}] 
	}
	sum := 0
	for _, val := range piles{
		sum += val
	}
	return maxStoneFrom(0,len(piles)-1).val > sum - maxStoneFrom(0,len(piles)-1).val
}
