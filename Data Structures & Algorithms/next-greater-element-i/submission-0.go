func nextGreaterElement(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	for i := 0; i < len(nums1); i++ {
		startProcess := false
		greaterVal := nums1[i]
		for j := 0; j < len(nums2); j++{
			if nums2[j] == nums1[i]{
				startProcess = true
				continue
			}
			if startProcess {
				if nums2[j] > greaterVal {
					greaterVal = nums2[j]
					break
				}
			}
		}
		if greaterVal == nums1[i]{
			result = append(result, -1)
		} else {
			result = append(result, greaterVal)
		}
	}
	return result
}
