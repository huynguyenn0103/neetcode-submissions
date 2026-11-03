
type Node struct {
    key   int
    value int
    freq  int
    prev  *Node
    next  *Node
}

type DoublyLinkedList struct {
    head *Node
    tail *Node
    size int
}

func NewList() *DoublyLinkedList {
    head := &Node{}
    tail := &Node{}

    head.next = tail
    tail.prev = head

    return &DoublyLinkedList{
        head: head,
        tail: tail,
    }
}

// Insert node at the head (most recently used)
func (l *DoublyLinkedList) AddFront(node *Node) {
    first := l.head.next

    node.prev = l.head
    node.next = first

    l.head.next = node
    first.prev = node

    l.size++
}

// Remove a specific node
func (l *DoublyLinkedList) Remove(node *Node) {
    node.prev.next = node.next
    node.next.prev = node.prev

    l.size--
}

// Remove the least recently used node
func (l *DoublyLinkedList) RemoveLast() *Node {
    if l.size == 0 {
        return nil
    }

    node := l.tail.prev
    l.Remove(node)

    return node
}

type LFUCache struct {
    capacity  int
    minFreq   int
    keyToNode map[int]*Node
    freqToList map[int]*DoublyLinkedList
}

func Constructor(capacity int) LFUCache {
    return LFUCache{
        capacity:   capacity,
        minFreq:    0,
        keyToNode:  make(map[int]*Node),
        freqToList: make(map[int]*DoublyLinkedList),
    }
}

// Increase the frequency of an existing node
func (c *LFUCache) UpdateFreq(node *Node) {
    oldFreq := node.freq

    // Remove from the old frequency list
    oldList := c.freqToList[oldFreq]
    oldList.Remove(node)

    // If this was the last node with minFreq,
    // the minimum frequency must increase.
    if oldFreq == c.minFreq && oldList.size == 0 {
        c.minFreq++
    }

    // Increase frequency
    node.freq++

    // Add to the new frequency list
    if c.freqToList[node.freq] == nil {
        c.freqToList[node.freq] = NewList()
    }

    c.freqToList[node.freq].AddFront(node)
}

func (c *LFUCache) Get(key int) int {
    node, exists := c.keyToNode[key]

    if !exists {
        return -1
    }

    // Accessing a key increases its frequency
    c.UpdateFreq(node)

    return node.value
}

func (c *LFUCache) Put(key int, value int) {
    if c.capacity == 0 {
        return
    }

    // Case 1: Key already exists
    if node, exists := c.keyToNode[key]; exists {
        node.value = value

        // Updating an existing key also increases frequency
        c.UpdateFreq(node)
        return
    }

    // Case 2: Cache is full -> evict LFU + LRU
    if len(c.keyToNode) == c.capacity {
        list := c.freqToList[c.minFreq]

        // Tail is the least recently used node
        removed := list.RemoveLast()

        delete(c.keyToNode, removed.key)
    }

    // Case 3: Insert a new key
    node := &Node{
        key:   key,
        value: value,
        freq:  1,
    }

    c.keyToNode[key] = node

    if c.freqToList[1] == nil {
        c.freqToList[1] = NewList()
    }

    c.freqToList[1].AddFront(node)

    // Every newly inserted key starts with frequency 1
    c.minFreq = 1
}
