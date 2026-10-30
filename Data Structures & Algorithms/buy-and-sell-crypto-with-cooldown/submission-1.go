func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }

    hold := -prices[0]
    sold := 0
    rest := 0

    for i := 1; i < len(prices); i++ {
        prevHold := hold
        prevSold := sold
        prevRest := rest

        hold = max(prevHold, prevRest-prices[i])
        sold = prevHold + prices[i]
        rest = max(prevRest, prevSold)
    }

    return max(sold, rest)
}