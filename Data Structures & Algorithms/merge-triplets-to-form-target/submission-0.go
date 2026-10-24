const Infinity = math.MaxInt

func mergeTriplets(triplets [][]int, target []int) bool {
    candidate := []int{-Infinity, -Infinity, -Infinity}
	for _, triplet := range triplets {
		x := max(triplet[0], candidate[0])
		y := max(triplet[1], candidate[1])
		z := max(triplet[2], candidate[2])
		if x <= target[0] && y <= target[1] && z <= target[2]{
			candidate = []int{x,y,z}
		}
	}
	return candidate[0] == target[0] && candidate[1] == target[1] && candidate[2] == target[2]
}
