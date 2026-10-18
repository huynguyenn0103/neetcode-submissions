

func findDuplicate(nums []int) int {
    slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	meet := slow
	slow = 0
	for slow != meet {
		meet = nums[meet]
		slow = nums[slow]
	}
	return slow
}
