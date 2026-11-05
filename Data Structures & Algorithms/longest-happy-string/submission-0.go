type HeapItem struct {
	c byte
	freq int
}

type MaxHeap []HeapItem
func (h MaxHeap)Len()int{return len(h)}
func (h MaxHeap)Less(i,j int)bool{
	if h[i].freq == h[j].freq{
		return h[i].c < h[j].c
	}
	return h[i].freq > h[j].freq
}
func (h MaxHeap)Swap(i,j int){h[i],h[j]=h[j],h[i]}
func (h *MaxHeap) Push(x interface{}){*h = append(*h, x.(HeapItem))}
func (h *MaxHeap) Pop()interface{}{
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

func longestDiverseString(a int, b int, c int) string {
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	if a > 0{
		heap.Push(maxHeap, HeapItem{c: 'a', freq: a})
	}
	if b > 0 {
		heap.Push(maxHeap, HeapItem{c: 'b', freq: b})
	}
	if c > 0{
		heap.Push(maxHeap, HeapItem{c: 'c', freq: c})
	}	
	result := make([]byte,0)
	for maxHeap.Len()>0{
		first := heap.Pop(maxHeap).(HeapItem)
		result = append(result, first.c)
		if first.freq > 1 {
			heap.Push(maxHeap, HeapItem{c: first.c, freq: first.freq-1})
		}
		if maxHeap.Len() == 0 {break}
		second := heap.Pop(maxHeap).(HeapItem)
		result = append(result, second.c)
		if second.c == first.c {
			if maxHeap.Len() == 0 {
				break
			} else {
				third := heap.Pop(maxHeap).(HeapItem)
				result = append(result,third.c)
				if third.freq > 1 {
					heap.Push(maxHeap, HeapItem{c: third.c, freq: third.freq-1})
				}
			}
		}
		if second.freq > 1 {
			heap.Push(maxHeap, HeapItem{c: second.c, freq: second.freq-1})
		}
	}
	return string(result)
}
