func minSubarray(nums []int, p int) int {
	sum := 0
	for _, num := range nums{
		sum = sum + num
	}
	if sum % p == 0 {
		return 0
	}

	targetRemainder := sum % p
	prefixSum := 0
	result := math.MaxInt

	mapRemainderToLatestIdx := make(map[int]int)
	//Find 2 prefix sum that minus have remainder = target
	for i, num := range nums{
		prefixSum += num
		remainder := prefixSum % p
		if remainder == targetRemainder && i != len(nums)-1 {
			result = min(result, i+1)
		}
		neededRemainder := remainder - targetRemainder
		if neededRemainder < 0 {
			neededRemainder += p
		}
		if idx, existed := mapRemainderToLatestIdx[neededRemainder]; existed{
			result = min(result, i - idx)
		}
		mapRemainderToLatestIdx[remainder] = i
	}
	if result == math.MaxInt{
		return -1
	}
	return result
}