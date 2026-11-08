import "slices"

func nextStr(str string) []string{
	bytes := []byte(str)
	result := make([]string,0)
	for i := 0; i < len(bytes); i++ {
		clone1 := slices.Clone(bytes)
		if clone1[i] == '9' {
			clone1[i] = '0'
		} else {
			clone1[i] = clone1[i] + 1
		}
		result = append(result, string(clone1))

		clone2 := slices.Clone(bytes)
		if clone2[i] == '0'{
			clone2[i] = '9'
		} else {
			clone2[i] = clone2[i] - 1
		}
		result = append(result, string(clone2))
	}
	return result
}

type Item struct {
	level int
	val string
}

func openLock(deadends []string, target string) int {
	visited := make(map[string]bool)
	for _, deadend := range deadends{
		visited[deadend]=true
	}
	queue := make([]Item,0)
	queue = append(queue, Item{level: 0, val: "0000"})
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		if visited[item.val]{continue}
		visited[item.val]=true
		if item.val == target{return item.level}
		nextVals := nextStr(item.val)
		for _, nextVal := range nextVals{
			queue = append(queue, Item{val: nextVal, level: item.level+1})
		}
	}
	return -1
}
