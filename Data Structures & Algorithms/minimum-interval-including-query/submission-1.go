

type Interval struct {
    length int
    right  int
}

type MinHeap []Interval

func (h MinHeap) Len() int {
    return len(h)
}

func (h MinHeap) Less(i, j int) bool {
    return h[i].length < h[j].length
}

func (h MinHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
    *h = append(*h, x.(Interval))
}

func (h *MinHeap) Pop() any {
    old := *h
    n := len(old)
    item := old[n-1]
    *h = old[:n-1]
    return item
}

func minInterval(intervals [][]int, queries []int) []int {
    // Sort intervals by left endpoint
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    // Preserve original query indices
    type Query struct {
        value int
        index int
    }

    sortedQueries := make([]Query, len(queries))

    for i, q := range queries {
        sortedQueries[i] = Query{value: q, index: i}
    }

    sort.Slice(sortedQueries, func(i, j int) bool {
        return sortedQueries[i].value < sortedQueries[j].value
    })

    result := make([]int, len(queries))

    h := &MinHeap{}
    heap.Init(h)

    i := 0

    for _, query := range sortedQueries {
        q := query.value

        // Add intervals that have started
        for i < len(intervals) && intervals[i][0] <= q {
            left := intervals[i][0]
            right := intervals[i][1]

            heap.Push(h, Interval{
                length: right - left + 1,
                right:  right,
            })

            i++
        }

        // Remove expired intervals from heap top
        for h.Len() > 0 && (*h)[0].right < q {
            heap.Pop(h)
        }

        if h.Len() == 0 {
            result[query.index] = -1
        } else {
            result[query.index] = (*h)[0].length
        }
    }

    return result
}