type MyStack struct {
	queue1 []int
	queue2 []int
}

func Constructor() MyStack {
	return MyStack{
		queue1 : make([]int,0),
		queue2 : make([]int,0),
	}
}

func (this *MyStack) Push(x int) {
	this.queue1 = append(this.queue1,x)
}

func (this *MyStack) Pop() int {
	for len(this.queue1) > 1{
		this.queue2 = append(this.queue2,this.queue1[0])
		this.queue1 = this.queue1[1:]
	}
	item := this.queue1[0]
	this.queue1 = this.queue2
	this.queue2 = make([]int,0)
	return item
}

func (this *MyStack) Top() int {
	item := this.Pop()
	this.queue1 = append(this.queue1,item)
	return item
}

func (this *MyStack) Empty() bool {
	return len(this.queue1) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Top();
 * param4 := obj.Empty();
 */
