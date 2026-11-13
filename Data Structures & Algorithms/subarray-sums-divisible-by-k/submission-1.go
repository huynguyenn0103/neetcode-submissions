func subarraysDivByK(nums []int, k int) int {
    result := 0
    prefixSum := 0

    remainderCount := make(map[int]int)
    remainderCount[0] = 1

    for _, num := range nums {
        prefixSum += num

        remainder := prefixSum % k

        // Handle negative numbers.
        if remainder < 0 {
            remainder += k
        }

        result += remainderCount[remainder]

        remainderCount[remainder]++
    }

    return result
}