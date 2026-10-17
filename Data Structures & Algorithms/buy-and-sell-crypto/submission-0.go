
func maxProfit(prices []int) int {
	n := len(prices)
	ptr1, ptr2 := 0, 1
	result := 0
	for ptr2 < n {
		for ptr2 < n && prices[ptr2] >= prices[ptr1]{
			result = max(result, prices[ptr2]-prices[ptr1])
			ptr2++
		}
		ptr1 = ptr2
		ptr2++
	}
	return result
}
