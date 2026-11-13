import "slices"

type Item struct{
	char byte
	idx int
}

func customSortString(order string, s string) string {
	mapCharToIdx := make(map[byte]int)
	for i := 0; i < len(order); i++{
		mapCharToIdx[order[i]] = i
	}
	items := make([]Item, 0)
	for i := 0; i < len(s); i++{
		items = append(items, Item{char: s[i], idx: mapCharToIdx[s[i]]})
	}
	slices.SortFunc(items, func(i1, i2 Item)int{
		return i1.idx - i2.idx
	})
	resultInBytes := make([]byte,0)
	for _, item := range items{
		resultInBytes = append(resultInBytes, item.char)
	}
	return string(resultInBytes)
}
