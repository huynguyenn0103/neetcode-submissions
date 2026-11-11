func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed); i++{
		prev, next := i-1, i+1
		if (prev >= 0 && flowerbed[prev] == 1) || 
		(next < len(flowerbed) && flowerbed[next] == 1) ||
		flowerbed[i] == 1{
			continue
		}
		flowerbed[i] = 1
		n--
	}
	return n <= 0
}