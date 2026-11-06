func maxTurbulenceSize(arr []int) int {
	result := 1
	n := len(arr)
	ptr1 := 0
	for ptr1 < n-1 {
		ptr2 := ptr1+1
		if arr[ptr1] != arr[ptr2]{
			startCompare := 1
			count := 2
			if arr[ptr1] > arr[ptr2]{
				startCompare = -1
			}
			for ptr2<n-1{
				if (startCompare == -1 && arr[ptr2] < arr[ptr2+1]) || (startCompare == 1 && arr[ptr2] > arr[ptr2+1]) {
					count++
					startCompare = -startCompare
				} else {
					break
				}
				ptr2++
			}
			result = max(result,count)
		}	
		ptr1=ptr2
	}
	return result
}
