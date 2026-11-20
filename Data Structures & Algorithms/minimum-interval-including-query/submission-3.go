
import (
    "slices"
)

type Interval struct {
    right  int
    length int
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

func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(Interval))
}

func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    item := old[n-1]
    *h = old[:n-1]
    return item
}

func minInterval(intervals [][]int, queries []int) []int {
    slices.SortFunc(intervals, func(a, b []int) int {
        return a[0] - b[0]
    })

    type Query struct {
        value int
        index int
    }

    sortedQueries := make([]Query, len(queries))

    for i, q := range queries {
        sortedQueries[i] = Query{
            value: q,
            index: i,
        }
    }

    slices.SortFunc(sortedQueries, func(a, b Query) int {
        return a.value - b.value
    })

    result := make([]int, len(queries))

    minHeap := &MinHeap{}
    heap.Init(minHeap)

    i := 0

    for _, q := range sortedQueries {
        // Add all intervals that have started.
        for i < len(intervals) && intervals[i][0] <= q.value {
            left, right := intervals[i][0], intervals[i][1]

            heap.Push(minHeap, Interval{
                right:  right,
                length: right - left + 1,
            })

            i++
        }

        // Remove intervals that have already ended.
        for minHeap.Len() > 0 && (*minHeap)[0].right < q.value {
            heap.Pop(minHeap)
        }

        // Get the shortest valid interval.
        result[q.index] = -1

        if minHeap.Len() > 0 {
            result[q.index] = (*minHeap)[0].length
        }
    }

    return result
}
