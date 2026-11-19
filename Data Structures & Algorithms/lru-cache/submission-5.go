type ListNode struct{
	val int
	prev *ListNode
	next *ListNode
}

func (l *ListNode) Unlink(){
	if l != nil {
		prev := l.prev
		next := l.next
		if prev != nil {
			prev.next = next
		}
		if next != nil {
			next.prev = prev
		}
		l.prev = nil
		l.next = nil
	}
}

type LinkedList struct{
	head *ListNode
	tail *ListNode
}

func NewLinkedList() LinkedList{
	return LinkedList{}
}

func (l *LinkedList) RemoveHead() *ListNode{
	if l.head == nil {return nil}
	item := l.head
	if l.head == l.tail {
		l.head = nil
		l.tail = nil 
	} else {
		l.head = l.head.next
		if l.head != nil {
			l.head.prev = nil
		}
	}
	item.next = nil
	return item
}

func (l *LinkedList) RemovedNode(node *ListNode){
	if node == nil {return}
	if l.head == node {
		l.head = l.head.next
	}
	if l.tail == node {
		l.tail = l.tail.prev
	}
	node.Unlink()
}

func (l *LinkedList) AppendNode(node *ListNode) {
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		node.prev = l.tail
		l.tail = l.tail.next
	}
}


type LRUCache struct {
	mapKeyToNode map[int]*ListNode
	mapNodeToKey map[*ListNode]int
    linkedList LinkedList
	curSize int
	capacity int
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
		linkedList: NewLinkedList(),
		curSize: 0,
		capacity: capacity,
		mapKeyToNode: make(map[int]*ListNode),
		mapNodeToKey: make(map[*ListNode]int),
	}
}

func (this *LRUCache) Get(key int) int {
    node, existed := this.mapKeyToNode[key]
	if !existed {
		return -1
	} else {
		this.linkedList.RemovedNode(node)
		this.linkedList.AppendNode(node)
		return node.val
	}
}

func (this *LRUCache) Put(key int, value int) {
    node, existed := this.mapKeyToNode[key]
	if existed{
		node.val = value
		this.linkedList.RemovedNode(node)
		this.linkedList.AppendNode(node)
	} else {
		newNode := &ListNode{val: value}
		if this.curSize == this.capacity{
			removed := this.linkedList.RemoveHead()
			removedKey := this.mapNodeToKey[removed]
			delete(this.mapNodeToKey, removed)
			delete(this.mapKeyToNode, removedKey)
			this.curSize--
		}
		this.mapKeyToNode[key] = newNode
		this.mapNodeToKey[newNode] = key
		this.curSize++
		this.linkedList.AppendNode(newNode)
	}
}
