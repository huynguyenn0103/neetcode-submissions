// 2,3 => 3,2 => 2,1 => 1,1

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func interchangeableRectangles(rectangles [][]int) int64 {
	mapPairToFreq := make(map[[2]int]int)
	result := int64(0)
	for _, rectangle := range rectangles{
		w,h := rectangle[0], rectangle[1]
		x := gcd(w,h)
		w,h = w/x,h/x
		result += int64(mapPairToFreq[[2]int{w,h}])
		mapPairToFreq[[2]int{w,h}]++
	}
	return result
}
