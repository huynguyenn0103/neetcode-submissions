
type Pair struct {
    char byte
    freq int
}

type MaxHeap []Pair

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
    return h[i].freq > h[j].freq
}
func (h MaxHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}
func (h *MaxHeap) Push(x interface{}) {
    *h = append(*h, x.(Pair))
}
func (h *MaxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    item := old[n-1]
    *h = old[:n-1]
    return item
}

func reorganizeString(s string) string {
    freq := make([]int, 26)

    for i := 0; i < len(s); i++ {
        freq[s[i]-'a']++
    }

    // Impossible if one character occurs too frequently.
    maxFreq := 0
    for _, f := range freq {
        if f > maxFreq {
            maxFreq = f
        }
    }

    if maxFreq > (len(s)+1)/2 {
        return ""
    }

    h := &MaxHeap{}
    heap.Init(h)

    for i, f := range freq {
        if f > 0 {
            heap.Push(h, Pair{byte(i) + 'a', f})
        }
    }

    result := make([]byte, 0, len(s))

    for h.Len() > 0 {
        first := heap.Pop(h).(Pair)
        result = append(result, first.char)
        first.freq--

        if h.Len() == 0 {
            break
        }

        second := heap.Pop(h).(Pair)
        result = append(result, second.char)
        second.freq--

        if first.freq > 0 {
            heap.Push(h, first)
        }

        if second.freq > 0 {
            heap.Push(h, second)
        }
    }

    return string(result)
}