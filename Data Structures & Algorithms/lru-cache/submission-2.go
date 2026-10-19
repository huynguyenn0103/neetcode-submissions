type ListNode struct {
	prev *ListNode
	next *ListNode
	val int
	key int
}

type LinkedList struct {
	head *ListNode
	tail *ListNode
}

// Put + not existed + not full
func (l *LinkedList) AppendHead(target *ListNode){
	if l.head == nil {
		l.head = target
		l.tail = target
	} else {
		target.next = l.head
		l.head.prev = target
		l.head = target
	}
}
// Put + not existed + full
func (l *LinkedList) RemoveLast() *ListNode{
	item := l.tail
	if item.prev != nil {
		item.prev.next = nil
		l.tail = item.prev
	} else {
		l.head = nil
		l.tail = nil
	}
	item.prev = nil
	return item
}

// Put + Existed
func (l *LinkedList) MoveNodeToHead(target *ListNode){
	if target.next != nil {
		target.next.prev = target.prev
	}
	if target.prev != nil {
		target.prev.next = target.next
	}
	if target == l.head {
		l.head = target.next
	}
	if target == l.tail {
		l.tail = target.prev
	}
	target.next = nil
	target.prev = nil
	l.AppendHead(target)
}

type LRUCache struct {
    capacity int
	linkedList LinkedList
	mapKeyToNode map[int]*ListNode
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
		capacity: capacity,
		linkedList: LinkedList{head: nil, tail: nil},
		mapKeyToNode: make(map[int]*ListNode),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, existed := this.mapKeyToNode[key]; !existed{
		return -1
	} else {
		this.linkedList.MoveNodeToHead(node)
		return node.val
	}
}

func (this *LRUCache) Put(key int, value int) {
    node, existed := this.mapKeyToNode[key]
	if existed {
		node.val = value
		this.linkedList.MoveNodeToHead(node)
	} else {
		curLen := len(this.mapKeyToNode)
		newNode := &ListNode{val: value, key: key}
		if curLen == this.capacity{
			removeNode := this.linkedList.RemoveLast()
			delete(this.mapKeyToNode, removeNode.key)
		}
		this.mapKeyToNode[key]=newNode
		this.linkedList.AppendHead(newNode)
	}
}
