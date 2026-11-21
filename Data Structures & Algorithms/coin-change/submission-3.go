type Result struct{
	fewest int
	isOK bool
}

func coinChange(coins []int, amount int) int {
    cache := make(map[int]Result)
	var fewestCoin func(amount int)(int,bool)
	fewestCoin = func(amount int)(int,bool){
		if amount < 0 {
			return 0, false
		}
		if amount == 0 {
			return 0, true
		}
		if val, existed := cache[amount]; existed{
			return val.fewest, val.isOK
		}
		result := math.MaxInt
		for _, coin := range coins{
			fewest, isOk := fewestCoin(amount-coin)
			if !isOk{continue}
			result = min(result, fewest+1)
		}
		var val Result
		if result != math.MaxInt{
			val.isOK = true
			val.fewest = result
		}
		cache[amount]=val
		return val.fewest, val.isOK
	}
	result, isOK := fewestCoin(amount)
	if !isOK{
		return -1
	}
	return result
}
