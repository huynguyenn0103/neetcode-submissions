import "slices"
const Infinity = math.MaxInt
func coinChange(coins []int, amount int) int {
	slices.SortFunc(coins, func(i1,i2 int)int{return i1-i2})
	minVal := Infinity
	var recursion func(coinIdx, remain, total int)
	recursion = func(coinIdx, remain, total int){
		if remain == 0 {
			minVal = min(minVal, total)
			return
		}
		if coinIdx < 0 {return}
		maxPossible := remain / coins[coinIdx]
		for i := maxPossible; i >= 0; i--{
			recursion(coinIdx-1, remain - i*coins[coinIdx], total + i)
		}
	}
	recursion(len(coins)-1,amount,0)
	if minVal == Infinity{return -1}
	return minVal
}
