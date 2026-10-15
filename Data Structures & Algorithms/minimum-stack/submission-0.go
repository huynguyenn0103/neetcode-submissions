type MinStack struct {
	minStack []int
	stack []int
}

func Constructor() MinStack {
	return MinStack{
		minStack: make([]int, 0),
		stack: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.minStack) == 0 || val <= this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, val)
	}
}

func (this *MinStack) Pop() {
	valToRemove := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if len(this.minStack) > 0 {
		topMinStack := this.minStack[len(this.minStack)-1]
		if topMinStack == valToRemove{
			this.minStack = this.minStack[:len(this.minStack)-1]
		}
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
