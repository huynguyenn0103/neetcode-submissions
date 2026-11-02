func rotate(nums []int, k int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	// when k > n
	k = k % n

	// reverse array 
	reverse(nums, 0 , n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, l, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}