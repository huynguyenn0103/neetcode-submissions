// 1 2 2 3 3
import "slices"
func numRescueBoats(people []int, limit int) int {
	slices.SortFunc(people, func(i1,i2 int)int{return i1-i2})
	n := len(people)
	left, right := 0, n-1
	result := 0
	curLimit := limit

	for right >= 0 && people[right] == limit {
		right--
		result++
	}
	numOfPeople := 0
	resetAndMove := func () {
		result++
		curLimit = limit
		numOfPeople = 0
	}
	for left <= right {
		if curLimit >= people[right] {
			curLimit = curLimit - people[right]
			numOfPeople++
			if curLimit == 0 || numOfPeople == 2 {
				resetAndMove()
			}
			right--
		} else if curLimit >= people[left] {
			curLimit = curLimit - people[left]
			numOfPeople++
			if curLimit == 0 || numOfPeople == 2 {
				resetAndMove()
			}
			left++
		} else {
			resetAndMove()
		}
	}
	if curLimit != limit {
		result++
	}
	return result
}
