import "slices"

func change(amount int, coins []int) int {
	slices.Sort(coins)
    cacheNumOfWayToBuildAmount := make(map[[2]int]int)
	var numOfWayToBuildAmount func(amount, minCoin int)int
	numOfWayToBuildAmount = func(amount, minCoin int)int{
		if amount == 0 {return 1}
		if amount < 0 {
			return 0
		} 
		key := [2]int{amount,minCoin}
		if val, existed := cacheNumOfWayToBuildAmount[key]; existed{
			return val
		}
		result := 0
		for _, coin := range coins{
			if coin < minCoin{continue}
			result += numOfWayToBuildAmount(amount-coin,coin)
		}
		cacheNumOfWayToBuildAmount[key]=result
		return result
	}
	return numOfWayToBuildAmount(amount,0)
}
