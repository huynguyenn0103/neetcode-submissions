
type Task struct {
    freq int
    nextAvailable int
}

type MaxHeap []Task

func (h MaxHeap) Len() int {
    return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
    return h[i].freq > h[j].freq
}

func (h MaxHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

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
    freq := make(map[byte]int)
    for _, task := range tasks {
        freq[task]++
    }

    maxHeap := &MaxHeap{}
    heap.Init(maxHeap)

    for _, count := range freq {
        heap.Push(maxHeap, Task{freq: count})
    }

    cooldown := make([]Task, 0)
    time := 0

    for maxHeap.Len() > 0 || len(cooldown) > 0 {
        time++

        // Release tasks whose cooldown has expired.
        for i := 0; i < len(cooldown); {
            if cooldown[i].nextAvailable <= time {
                heap.Push(maxHeap, cooldown[i])
                cooldown = append(cooldown[:i], cooldown[i+1:]...)
            } else {
                i++
            }
        }

        if maxHeap.Len() == 0 {
            continue
        }

        // Execute the task with the highest remaining frequency.
        task := heap.Pop(maxHeap).(Task)
        task.freq--

        if task.freq > 0 {
            task.nextAvailable = time + n + 1
            cooldown = append(cooldown, task)
        }
    }

    return time
}