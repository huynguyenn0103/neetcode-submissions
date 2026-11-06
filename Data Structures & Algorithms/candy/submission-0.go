// 5 4 3 2 1 2 3 4 5 6

func candy(ratings []int) int {
	if len(ratings) == 1{
		return 1
	}
	localMinIdx := make(map[int]bool,0)
	n := len(ratings)
	for i := 0; i < n; i++ {
		if (i == 0 && ratings[i] <= ratings[i+1]) ||
		(i == n-1 && ratings[i-1] >= ratings[i]) ||
		(i > 0 && i < n-1 && ratings[i] <= ratings[i-1] && ratings[i] <= ratings[i+1]){
			localMinIdx[i]=true
		}
	}
	result := make([]int,n)
	
	for localMin := range localMinIdx{
		count := 1
		left := localMin
		for left >= 0 {
			result[left] = max(result[left],count)
			left--
			if left < n-1 && left >= 0 {
				if ratings[left] < ratings[left+1] || localMinIdx[left] {
					break
				} else if ratings[left] > ratings[left+1] {
					count++
				}
			}
		}
		count = 1
		right := localMin
		for right < n {
			result[right] = max(result[right],count)
			right++
			if right > 0 && right < n {
				if ratings[right-1] > ratings[right] || localMinIdx[right]{
					break
				} else if ratings[right-1] < ratings[right]{
					count++
				}
			}
		}
	}
	sum := 0
	for _, val := range result{
		sum += val
	}
	return sum
}
