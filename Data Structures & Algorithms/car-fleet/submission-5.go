

func carFleet(target int, position []int, speed []int) int {
    n := len(position)

    // Pair each car's position with its speed.
    cars := make([][2]int, n)

    for i := 0; i < n; i++ {
        cars[i] = [2]int{position[i], speed[i]}
    }

    // Sort by position in descending order.
    sort.Slice(cars, func(i, j int) bool {
        return cars[i][0] > cars[j][0]
    })

    stack := make([]float64, 0, n)

    for _, car := range cars {
        pos := car[0]
        spd := car[1]

        time := float64(target-pos) / float64(spd)

        // If this car catches the fleet ahead,
        // it joins that fleet.
        if len(stack) > 0 && time <= stack[len(stack)-1] {
            continue
        }

        // Otherwise, it forms a new fleet.
        stack = append(stack, time)
    }

    return len(stack)
}