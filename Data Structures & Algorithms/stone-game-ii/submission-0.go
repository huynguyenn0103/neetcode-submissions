type Result struct{
	val int
	stones int
}

func stoneGameII(piles []int)int{
	cache := make(map[[2]int]Result)
	n := len(piles)
	var maxStoneFromIdxWithM func(idx, m int)Result
	maxStoneFromIdxWithM = func(idx, m int)Result{
		if idx >= n {return Result{val: 0, stones: 0}}
		if val, existed := cache[[2]int{idx,m}]; existed{
			return val
		}
		var result Result
		sum := 0
		bestOfOpponentStone := math.MaxInt
		for i := idx; i < idx + 2*m && i < n; i++ {
			sum = sum + piles[i]
			curStones := i-idx+1
			mNext := max(curStones,m)
			opponent:= maxStoneFromIdxWithM(i+1,mNext)
			if opponent.val < bestOfOpponentStone{
				result = Result{val: sum + maxStoneFromIdxWithM(i+1+opponent.stones,max(opponent.stones,mNext)).val, stones: curStones}
				bestOfOpponentStone = opponent.val
			}
		}
		cache[[2]int{idx,m}] = result
		return result
	}
	return maxStoneFromIdxWithM(0,1).val
}
