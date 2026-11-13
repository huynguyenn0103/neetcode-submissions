// 0, 1, 3 => |0-xk| + |1-xk| + |3-xk|

func minOperations(boxes string) []int {
    hasBallIndices := make([]int, 0)
	n := len(boxes)
	total := 0
	for i := 0; i < n; i++ {
		if boxes[i] == '1'{
			hasBallIndices = append(hasBallIndices, i)
			total = total + i
		}
	}
	m := len(hasBallIndices)
	result := make([]int, 0)

	curSum := 0
	ptr1, ptr2 := 0, 0
	for ptr2 < n{
		if ptr1 < m && ptr2 > hasBallIndices[ptr1]{
			curSum = curSum + hasBallIndices[ptr1]
			ptr1++
		}
		numOfBefore := ptr1
		numOfAfter := m - ptr1
		result = append(result, (numOfBefore*ptr2-curSum) + (total-curSum-numOfAfter*ptr2))
		ptr2++
	}
	return result
}