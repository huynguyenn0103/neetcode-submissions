type MyQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor() MyQueue {
	return MyQueue{
		stack1: make([]int,0),
		stack2: make([]int,0),
	}
}

func (this *MyQueue) Push(x int) {
	this.stack1 = append(this.stack1,x)
}

func (this *MyQueue) Pop() int {
	for len(this.stack1) > 1 {
		this.stack2 = append(this.stack2, this.stack1[len(this.stack1)-1])
		this.stack1 = this.stack1[:len(this.stack1)-1]
	}
	item := this.stack1[len(this.stack1)-1]
	this.stack1 = this.stack1[:len(this.stack1)-1]
	for len(this.stack2) > 0 {
		this.stack1 = append(this.stack1, this.stack2[len(this.stack2)-1])
		this.stack2 = this.stack2[:len(this.stack2)-1]
	}
	return item
}

func (this *MyQueue) Peek() int {
		for len(this.stack1) > 1 {
		this.stack2 = append(this.stack2, this.stack1[len(this.stack1)-1])
		this.stack1 = this.stack1[:len(this.stack1)-1]
	}
	item := this.stack1[len(this.stack1)-1]
	for len(this.stack2) > 0 {
		this.stack1 = append(this.stack1, this.stack2[len(this.stack2)-1])
		this.stack2 = this.stack2[:len(this.stack2)-1]
	}
	return item
}

func (this *MyQueue) Empty() bool {
	return len(this.stack1) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Peek();
 * param4 := obj.Empty();
 */
