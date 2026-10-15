import "slices"

func sortedStr(str string)string {
	listOfBytes := []byte(str)
	slices.SortFunc(listOfBytes, func(x,y byte)int{return int(x) - int(y)})
	return string(listOfBytes)
}

func groupAnagrams(strs []string) [][]string {
	correspondStrs := make([]string, 0)
	for _, str := range strs{
		correspondStrs = append(correspondStrs, sortedStr(str))
	}
	mapSortedStrToOrigin := make(map[string][]string)
	for i, sortedStr := range correspondStrs {
		if mapSortedStrToOrigin[sortedStr] == nil {
			mapSortedStrToOrigin[sortedStr]=make([]string, 0)
		}
		mapSortedStrToOrigin[sortedStr] = append(mapSortedStrToOrigin[sortedStr], strs[i])
	}
	result := make([][]string, 0)
	for _, vals := range mapSortedStrToOrigin{
		result = append(result, vals)
	}
	return result
}
