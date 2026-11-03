
func splitArray(nums []int, k int) int {
    left, right := 0, 0

    for _, num := range nums {
        if num > left {
            left = num
        }
        right += num
    }

    var canSplit func(limit int) bool

    canSplit = func(limit int) bool {
        count := 1
        sum := 0

        for _, num := range nums {
            if sum+num > limit {
                count++
                sum = num
            } else {
                sum += num
            }

            if count > k {
                return false
            }
        }

        return true
    }

    for left < right {
        mid := left + (right-left)/2

        if canSplit(mid) {
            right = mid
        } else {
            left = mid + 1
        }
    }

    return left
}
