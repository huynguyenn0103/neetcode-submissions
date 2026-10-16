type Item struct {
	visited bool
	cons int
}


func longestConsecutive(nums []int) int {
	mapNumberToVisited := make(map[int]Item)
	for _, val := range nums {
		mapNumberToVisited[val]=Item{visited: false, cons: 1}
	}
	result := 0
	for _, val := range nums {
		item := mapNumberToVisited[val]
		if item.visited{continue}
		count := 1
		for {
			if v, existed := mapNumberToVisited[val + count]; !existed{
				break
			} else {
				if v.visited {
					count = count + v.cons
					break
				} else {
					count++
					mapNumberToVisited[val]=Item{visited: true, cons: v.cons}
				}
			}
		}
		mapNumberToVisited[val]=Item{visited: true, cons: count}
		result = max(result, count)
	}
	return result
}
