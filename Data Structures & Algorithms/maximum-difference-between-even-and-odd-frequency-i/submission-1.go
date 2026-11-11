func abs(x int)int{
	if x >= 0 {return x}
	return -x
}

func maxDifference(s string) int {
	mapByteToFreq := make(map[byte]int)
	n := len(s)
	for i := 0; i < n; i++ {
		mapByteToFreq[s[i]]++
	}
	result := -math.MaxInt
	for _, freq1 := range mapByteToFreq{
		for _, freq2 := range mapByteToFreq{
			if (freq1 % 2 == 0 && freq2 % 2 != 0) {
				result = max(result, freq2 - freq1)
			} else if (freq1 % 2 != 0 && freq2 % 2 == 0){
				result = max(result, freq1 - freq2)
			}
		}
	}
	return result
}
