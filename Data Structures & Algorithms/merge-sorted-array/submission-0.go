func merge(nums1 []int, m int, nums2 []int, n int) {
	result := make([]int, 0)
	ptr1,ptr2 := 0, 0
	for ptr1 < m && ptr2 < n {
		if nums1[ptr1] < nums2[ptr2]{
			result = append(result, nums1[ptr1])
			ptr1++
		} else {
			result = append(result, nums2[ptr2])
			ptr2++
		}
	}
	for ptr1 < m {
		result = append(result, nums1[ptr1])
		ptr1++
	}
	for ptr2 < n {
		result = append(result, nums2[ptr2])
		ptr2++
	}
	for i := range result {
		nums1[i]=result[i]
	}
}
