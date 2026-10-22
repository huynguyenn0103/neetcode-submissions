type CountSquares struct {
    points map[[2]int]int
}

func Constructor() CountSquares {
    return CountSquares{
        points: make(map[[2]int]int),
    }
}

func (this *CountSquares) Add(point []int) {
    key := [2]int{point[0], point[1]}
    this.points[key]++
}

func (this *CountSquares) Count(point []int) int {
    x, y := point[0], point[1]
    count := 0

    for diagonal, freq := range this.points {
        dx, dy := diagonal[0], diagonal[1]

        // A valid diagonal must form a square.
        if dx == x || abs(dx-x) != abs(dy-y) {
            continue
        }

        // The other two corners of the square.
        p1 := [2]int{x, dy}
        p2 := [2]int{dx, y}

        count += freq * this.points[p1] * this.points[p2]
    }

    return count
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}