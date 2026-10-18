func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    // Always binary search on the smaller array.
    if len(nums1) > len(nums2) {
        nums1, nums2 = nums2, nums1
    }

    m, n := len(nums1), len(nums2)

    leftSize := (m + n + 1) / 2

    left, right := 0, m

    for left <= right {
        i := left + (right-left)/2
        j := leftSize - i

        left1, right1 := math.MinInt, math.MaxInt
        left2, right2 := math.MinInt, math.MaxInt

        if i > 0 {
            left1 = nums1[i-1]
        }

        if i < m {
            right1 = nums1[i]
        }

        if j > 0 {
            left2 = nums2[j-1]
        }

        if j < n {
            right2 = nums2[j]
        }

        // Correct partition.
        if left1 <= right2 && left2 <= right1 {
            if (m+n)%2 == 1 {
                return float64(max(left1, left2))
            }

            return (float64(max(left1, left2)) +
                float64(min(right1, right2))) / 2
        }

        if left1 > right2 {
            // Too many elements from nums1 on the left.
            right = i - 1
        } else {
            // Too few elements from nums1 on the left.
            left = i + 1
        }
    }

    return 0
}