const mod = int(1e9 + 7)

func numOfSubarrays(arr []int) int {
    numOfEvenSum := 1
	numOfOddSum := 0
	result := 0
	prefix := 0
	n := len(arr)
	for i := 0; i < n; i++{
		prefix = prefix + arr[i]
		if prefix % 2 == 0 {
			result = (result + numOfOddSum) % mod
			numOfEvenSum++
		} else {
			result = result + numOfEvenSum % mod
			numOfOddSum++
		}
	}
	return result
}