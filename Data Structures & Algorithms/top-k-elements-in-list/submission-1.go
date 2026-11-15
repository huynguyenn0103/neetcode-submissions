import "slices"

type Item struct{
	num int
	freq int
}

func topKFrequent(nums []int, k int) []int {
	mapNumToFreq := make(map[int]int)
	for _, num := range nums{
		mapNumToFreq[num]++
	}
	items := make([]Item,0)
	for num, freq := range mapNumToFreq{
		items = append(items, Item{num,freq})
	}
	slices.SortFunc(items,func(i1,i2 Item)int{return i2.freq-i1.freq})
	result := make([]int,0)
	for k > 0 {
		item := items[0]
		items = items[1:]
		result = append(result, item.num)
		k--
	}
	return result
}
