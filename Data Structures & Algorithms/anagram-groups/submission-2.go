// O(l) & O(1)
func buildListsFrom(str string)[26]int{
	var freq [26]int
	n := len(str)
	for i := 0; i < n; i++{
		c := str[i]
		freq[int(c-'a')]++
	}
	return freq
}


// O (n * (l + n *l)) = O(n^2*l)
// O(n)

func groupAnagrams(strs []string) [][]string {
	mapGroupToStrs := make(map[[26]int][]string)
	result := make([][]string,0)
	for _, str := range strs{
		lists := buildListsFrom(str)
		if mapGroupToStrs[lists] == nil{
			mapGroupToStrs[lists] = make([]string,0)
		}
		mapGroupToStrs[lists] = append(mapGroupToStrs[lists],str)
	}
	for _, val := range mapGroupToStrs{
		result = append(result, val)
	}
	return result
}
