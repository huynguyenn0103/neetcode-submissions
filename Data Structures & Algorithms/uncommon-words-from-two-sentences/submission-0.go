func inMap1NotInMap2(map1, map2 map[string]int)[]string{
	result := make([]string, 0)
	for key, val := range map1{
		if val == 1 && map2[key] == 0{
			result = append(result, key)
		}
	}
	return result
}

func uncommonFromSentences(s1 string, s2 string) []string {
	mapS1 := make(map[string]int)
	mapS2 := make(map[string]int)
	strs1 := strings.Split(s1, " ")
	for _, str1 := range strs1{
		mapS1[str1]++
	}
	strs2 := strings.Split(s2, " ")
	for _, str2 := range strs2{
		mapS2[str2]++
	}
	result := make([]string, 0)
	result = append(result, inMap1NotInMap2(mapS1, mapS2)...)
	result = append(result, inMap1NotInMap2(mapS2, mapS1)...)
	return result
}
