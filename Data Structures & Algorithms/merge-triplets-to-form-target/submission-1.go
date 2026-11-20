

func mergeTriplets(triplets [][]int, target []int) bool {
    result := []int{0,0,0}
	for _, triplet := range triplets{
		if triplet[0]<=target[0] && triplet[1]<=target[1] && triplet[2] <= target[2]{
			result = []int{max(result[0],triplet[0]),max(result[1],triplet[1]),max(result[2],triplet[2])}
		}
	}
	return result[0]==target[0] && result[1]==target[1] && result[2] == target[2]
}
