func inMap1NotInMap2(map1, map2 map[int]bool)[]int{
	result := make([]int, 0)
	for key := range map1{
		if map2[key] == false{
			result = append(result, key)
		}
	}
	return result
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	mapNum1 := make(map[int]bool)
	mapNum2 := make(map[int]bool)
	for _, num := range nums1{
		mapNum1[num]=true
	}
	for _, num := range nums2{
		mapNum2[num]=true
	}
	result := make([][]int, 0)
	result = append(result, inMap1NotInMap2(mapNum1,mapNum2))
	result = append(result, inMap1NotInMap2(mapNum2,mapNum1))
	return result
}
