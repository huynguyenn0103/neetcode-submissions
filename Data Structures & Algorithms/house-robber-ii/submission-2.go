// 1 2 3 4 5 6 7 8 9 10

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0],nums[1])
	}
	if len(nums) == 3{
		return max(nums[0],nums[1],nums[2])
	}
    includeFirstIdx := make(map[int]bool)
	includeFirstIdx[0]=true
	n := len(nums)
	result := make([]int,n)
	result1 := make([]int, n)
	for i := 0; i < n; i++{
		if i <= 1 {
			result[i]=nums[i]
			if i == 0 {
				result1[i]=0
			} else {
				result1[i]=nums[i]
			}
		} else if i == 2{
			includeFirstIdx[i]=true
			result[i]=nums[i]+result[i-2]
			result1[i]=nums[i]+result1[i-2]
		} else {
			idx1, idx2 := i-2, i-3
			val1, val2 := result[idx1], result[idx2]
			if val1 > val2 {
				result[i] = nums[i] + val1
				if includeFirstIdx[idx1]{includeFirstIdx[i]=true}
			} else if val2 > val1{
				result[i] = nums[i] + val2
				if includeFirstIdx[idx2]{includeFirstIdx[i]=true}
			} else {
				result[i] = nums[i] + val1
				if includeFirstIdx[idx1] && includeFirstIdx[idx2]{
					includeFirstIdx[i]=true
				}
			}
			result1[i]=max(result1[i-2],result1[i-3])+nums[i]
		}
	}
	if result[n-2] >= result[n-1]{
		return result[n-2]
	} else {
		if !includeFirstIdx[n-1]{
			return result[n-1]
		}
		return max(result[n-2],result1[n-1],result[n-3])
	}

}
