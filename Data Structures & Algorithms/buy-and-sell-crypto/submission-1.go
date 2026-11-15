func maxProfit(prices []int) int {
	minFromStart := math.MaxInt
	profit := 0
	for _, price := range prices{
		if price < minFromStart {
			minFromStart = price
		} else {
			profit = max(profit, price - minFromStart)
		}
	}
	return profit
}
