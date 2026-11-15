// 1 2 2 3 3 4 4 5

// 1 2 3
// 2 3
import "slices"

func isNStraightHand(hand []int, groupSize int) bool {
    n := len(hand)
	if n % groupSize != 0 {return false}
	numOfGroups := n / groupSize
	groups := make([][]int, numOfGroups)
	slices.Sort(hand)
	for _, h := range hand{
		canPut := false
		for i, group := range groups{
			g := len(group)
			if g == groupSize{
				continue
			}
			if g == 0 || group[g-1] + 1 == h{
				canPut = true
				groups[i] = append(groups[i],h)
				break
			}
		}
		if !canPut{
			return false
		}
	}
	return true
}
