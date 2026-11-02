// 2 3 4 5 6 7 8 1

func rotate(nums []int, k int) {
	n := len(nums)
	for k > 0 {
		prev := nums[0]
		for i := 1; i < n; i++{
			tmp := nums[i]
			nums[i]=prev
			prev = tmp
		}
		nums[0]=prev
		k--
	}
}
