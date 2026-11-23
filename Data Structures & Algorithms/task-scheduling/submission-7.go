
type Task struct {
    name byte
    freq int
}

type MaxHeap []Task

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].freq > h[j].freq }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
    *h = append(*h, x.(Task))
}

func (h *MaxHeap) Pop() any {
    old := *h
    n := len(old)
    item := old[n-1]
    *h = old[:n-1]
    return item
}

func leastInterval(tasks []byte, n int) int {
    freq := [26]int{}
    for _, task := range tasks {
        freq[task-'A']++
    }

    maxHeap := &MaxHeap{}
    heap.Init(maxHeap)

    for i, count := range freq {
        if count > 0 {
            heap.Push(maxHeap, Task{
                name: byte(i) + 'A',
                freq: count,
            })
        }
    }

    nextAvailable := [26]int{}
    time := 0

    for maxHeap.Len() > 0 {
        temp := make([]Task, 0, maxHeap.Len())
        var selected Task
        found := false

        for maxHeap.Len() > 0 {
            task := heap.Pop(maxHeap).(Task)

            if nextAvailable[task.name-'A'] <= time {
                selected = task
                found = true
                break
            }

            temp = append(temp, task)
        }

        for _, task := range temp {
            heap.Push(maxHeap, task)
        }

        if !found {
            // Jump directly to the earliest available task.
            earliest := math.MaxInt
            for _, task := range temp {
                earliest = min(earliest, nextAvailable[task.name-'A'])
            }
            time = earliest
            continue
        }

        // Execute one task.
        time++
        selected.freq--

        if selected.freq > 0 {
            nextAvailable[selected.name-'A'] = time + n
            heap.Push(maxHeap, selected)
        }
    }

    return time
}