import "slices"

type Action struct{
	name string
	timestamp int
	website string
}

func isPatternSmaller(p1, p2 [3]string) bool {
    for i := 0; i < 3; i++ {
        if p1[i] < p2[i] {
            return true
        }
        if p1[i] > p2[i] {
            return false
        }
    }
    return false
}

func mostVisitedPattern(username []string, timestamp []int, website []string) []string {
	mapPatternToUsers := make(map[[3]string]int)
	actions := make([]Action, 0)
	n := len(username)
	for i := 0; i < n; i++ {
		actions = append(actions, Action{
			name: username[i],
			timestamp: timestamp[i],
			website: website[i],
		})
	}
	slices.SortFunc(actions, func(a1, a2 Action) int{return a1.timestamp - a2.timestamp})
	mapUserToWebsites := make(map[string][]string)
	for _, action := range actions{
		name, website := action.name, action.website
		if mapUserToWebsites[name] == nil{
			mapUserToWebsites[name] = make([]string, 0)
		}
		mapUserToWebsites[name] = append(mapUserToWebsites[name], website)
	}
	var result [3]string
	var count int
	for _, websites := range mapUserToWebsites{
		visited := make(map[[3]string]bool)
		m := len(websites)
		for i := 0; i < m; i++ {
			for j := i + 1; j < m; j++ {
				for k := j + 1; k < m; k++ {
					key := [3]string{websites[i],websites[j],websites[k]}
					if visited[key]{
						continue
					}
					visited[key]=true
					mapPatternToUsers[key]++
					result = key
					count = mapPatternToUsers[key]
				}
			}
		}
	}
	for key, val := range mapPatternToUsers{
		if val > count {
			count = val
			result  = key
		} else if val == count{
			if isPatternSmaller(key, result){
				result = key
			}
		}
	}
	return result[:]
}
