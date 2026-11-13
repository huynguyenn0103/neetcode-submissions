// 1 2 1 2 1 2 2

func minimumIndex(nums []int) int {
    var dominantElement int
	mapNumToFreq := make(map[int]int)
	n := len(nums)
	for _, num := range nums{
		mapNumToFreq[num]++
		if mapNumToFreq[num] > n/2{
			dominantElement = num
		}
	}
	countDominant := 0
	for i, num := range nums {
		if num == dominantElement{
			countDominant++
		}
		if countDominant > (i + 1)/2 && mapNumToFreq[dominantElement] - countDominant > (n - i - 1) / 2{
			return i
		}
	}
	return -1
}