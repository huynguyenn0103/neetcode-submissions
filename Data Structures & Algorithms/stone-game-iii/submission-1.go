type Result struct{
	val int
	stones int
}

func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)
	cache := make(map[int]Result)
	var highestScoreFromIdx func(idx int) Result
	highestScoreFromIdx = func(idx int) Result{
		if idx >= n {return Result{val: 0, stones: 0}}
		if val, existed := cache[idx]; existed{
			return val
		}
		myBest := -math.MaxInt
		opponentBest := math.MaxInt
		numStones := 0
		sum := 0
		for i := 1; i <= 3 && idx + i <= n; i++ {
			sum += stoneValue[idx + i - 1]
			bestChoiceOfOpponent := highestScoreFromIdx(idx + i)
			opponentStones := bestChoiceOfOpponent.stones
			bestChoiceOfMe := sum + (highestScoreFromIdx(idx + i + opponentStones)).val
			if bestChoiceOfOpponent.val < opponentBest && bestChoiceOfMe > myBest{
				myBest = bestChoiceOfMe
				opponentBest = bestChoiceOfOpponent.val
				numStones = i
			}
		}
		result := Result{val: myBest, stones: numStones}
		cache[idx]=result
		return cache[idx]
	}

	bestAlice := highestScoreFromIdx(0)

	sum := 0
	for _, val := range stoneValue {
		sum = sum + val
	}
	
	bob := sum - bestAlice.val
	if bob == bestAlice.val {
		return "Tie"
	} else if bob < bestAlice.val {
		return "Alice"
	} else {
		return "Bob"
	}
}
