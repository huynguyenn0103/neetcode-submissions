import "slices"

type Item struct {
	val int
	freq int
}

func topKFrequent(nums []int, k int) []int {
	mapValueToFreq := make(map[int]int)
	for _, val := range nums {
		mapValueToFreq[val]++
	}
	listOfItem := make([]Item, 0)
	for val, freq := range mapValueToFreq{
		listOfItem = append(listOfItem, Item{val, freq})
	}
	slices.SortFunc(listOfItem, func(x,y Item)int{return y.freq - x.freq})
	result := make([]int, 0)
	for _, item := range listOfItem[:k] {
		result = append(result, item.val)
	}
	return result
}
