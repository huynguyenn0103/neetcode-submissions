type StockSpanner struct {
	stack []int
}

func Constructor() StockSpanner {
	return StockSpanner {
		stack: make([]int,0),
	}
}

func (this *StockSpanner) Next(price int) int {
	day := 1
	tmpStack := make([]int,0)
	for len(this.stack) > 0 && this.stack[len(this.stack)-1] <= price {
		day++
		tmpStack = append(tmpStack, this.stack[len(this.stack)-1])
		this.stack = this.stack[:len(this.stack)-1]
	}
	for len(tmpStack) > 0 {
		this.stack = append(this.stack, tmpStack[len(tmpStack)-1])
		tmpStack = tmpStack[:len(tmpStack)-1]
	}
	this.stack = append(this.stack, price)
	return day
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor()
 * param1 := obj.Next(price)
 */
 