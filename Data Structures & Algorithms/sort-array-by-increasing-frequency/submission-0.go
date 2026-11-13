import "slices"

type Item struct{
	num int
	freq int
}

func frequencySort(nums []int) []int {
	mapNumToFreq := make(map[int]int)
	for _, num := range nums{
		mapNumToFreq[num]++
	}
	items := make([]Item, 0)
	for num, freq := range mapNumToFreq{
		items = append(items, Item{num: num, freq: freq})
	}
	slices.SortFunc(items, func(i1, i2 Item)int{
		if i1.freq == i2.freq{
			return i2.num - i1.num
		}
		return i1.freq - i2.freq
	})
	result := make([]int, 0)
	for _, item := range items{
		num, freq := item.num, item.freq
		for freq > 0 {
			result = append(result, num)
			freq--
		}
	}
	return result
}
