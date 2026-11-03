// [3,3]

type ListNode struct {
	val int
	next *ListNode
	prev *ListNode
}

type MyCircularQueue struct {
	head *ListNode
	tail *ListNode
	capacity int
	size int
}


func Constructor(k int) MyCircularQueue {
    return MyCircularQueue{
		capacity: k,
	}
}


func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.size == this.capacity {return false}
	this.size++
	newNode := &ListNode{val: value}
	if this.head == nil {
		this.head = newNode
		this.tail = newNode
	} else {
		newNode.prev = this.tail
		this.tail.next = newNode
		this.tail = newNode
	}
	return true
}


func (this *MyCircularQueue) DeQueue() bool {
    if this.size == 0 {return false}
	if this.size == 1 {
		this.head = nil
		this.tail = nil
	} else {
		tmp := this.head
		this.head = this.head.next
		this.head.prev = nil
		tmp.next = nil
	}
	this.size--
	return true
}


func (this *MyCircularQueue) Front() int {
    if this.head == nil {
		return -1
	} else {
		return this.head.val
	}
}


func (this *MyCircularQueue) Rear() int {
    if this.tail == nil {
		return -1
	} else {
		return this.tail.val
	}
}


func (this *MyCircularQueue) IsEmpty() bool {
    return this.size == 0
}


func (this *MyCircularQueue) IsFull() bool {
    return this.size == this.capacity
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param1 := obj.EnQueue(value);
 * param2 := obj.DeQueue();
 * param3 := obj.Front();
 * param4 := obj.Rear();
 * param5 := obj.IsEmpty();
 * param6 := obj.IsFull();
 */
 