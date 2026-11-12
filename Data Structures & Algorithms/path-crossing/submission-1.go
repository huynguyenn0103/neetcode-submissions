var mapActions = map[byte][2]int{
	'N': [2]int{-1,0},
	'E': [2]int{0,1},
	'S': [2]int{1,0},
	'W': [2]int{0,-1},
}

func isPathCrossing(path string) bool {
    mapVisited := make(map[[2]int]bool)
	mapVisited[[2]int{0,0}]=true
	cur := [2]int{0,0}
	n := len(path)
	for i:=0; i < n; i++{
		action := mapActions[path[i]]
		cur = [2]int{cur[0]+action[0], cur[1]+action[1]}
		if mapVisited[cur]{return true}
		mapVisited[cur]=true
	}
	return false
}