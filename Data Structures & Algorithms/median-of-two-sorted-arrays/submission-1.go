// 1 2 3 4 5
// 6 7 8 9

// 1 2 3 4  5
// 6        7 8 9
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    // Always binary search the smaller array.
    if len(nums1) > len(nums2) {
        return findMedianSortedArrays(nums2, nums1)
    }

    m, n := len(nums1), len(nums2)

    low, high := 0, m
    half := (m + n + 1) / 2

    for low <= high {
        i := (low + high)/2
        j := half - i

        // Boundary values of both partitions.
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

            return float64(max(left1, left2)+min(right1, right2)) / 2.0
        }

        // Too many elements from nums1.
        if left1 > right2 {
            high = i - 1
        } else {
            // Too few elements from nums1.
            low = i + 1
        }
    }

    return 0
}
