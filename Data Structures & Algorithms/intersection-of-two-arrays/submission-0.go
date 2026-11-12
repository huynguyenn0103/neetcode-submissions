func intersection(nums1 []int, nums2 []int) []int {
    mapResult := make(map[int]bool)
	mapNums1 := make(map[int]bool)
	for _, num1 := range nums1{
		mapNums1[num1]=true
	}
	for _, num2 := range nums2{
		if mapNums1[num2]{
			mapResult[num2]=true
		}
	}

	result := make([]int,0)
	for key := range mapResult{
		result = append(result, key)
	}
	return result
}
