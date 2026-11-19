type Item struct{
	val int
	minTilItem int
}

type MinStack struct {
	items []Item
}

func Constructor() MinStack {
	return MinStack{items: make([]Item,0)}
}

func (this *MinStack) Push(val int) {
	var minVal int
	if len(this.items) == 0{
		minVal = val
	} else {
		minVal = min(val, this.items[len(this.items)-1].minTilItem)
	}
	this.items = append(this.items, Item{val: val, minTilItem: minVal})
}

func (this *MinStack) Pop() {
	this.items = this.items[:len(this.items)-1]
}

func (this *MinStack) Top() int {
	return this.items[len(this.items)-1].val
}

func (this *MinStack) GetMin() int {
	return this.items[len(this.items)-1].minTilItem
}
