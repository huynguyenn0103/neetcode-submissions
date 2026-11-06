func predictPartyVictory(senate string) string {
    n := len(senate)

    rQueue := []int{}
    dQueue := []int{}

    for i := 0; i < n; i++ {
        if senate[i] == 'R' {
            rQueue = append(rQueue, i)
        } else {
            dQueue = append(dQueue, i)
        }
    }

    for len(rQueue) > 0 && len(dQueue) > 0 {
        r := rQueue[0]
        d := dQueue[0]

        rQueue = rQueue[1:]
        dQueue = dQueue[1:]

        if r < d {
            rQueue = append(rQueue, r+n)
        } else {
            dQueue = append(dQueue, d+n)
        }
    }

    if len(rQueue) > 0 {
        return "Radiant"
    }

    return "Dire"
}