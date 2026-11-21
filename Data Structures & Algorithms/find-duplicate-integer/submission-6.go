//145223
func findDuplicate(nums []int) int {
    slow := nums[0]
	fast := nums[nums[0]]
	for slow != fast {
		fast = nums[nums[fast]]
	    slow = nums[slow]
	}
	tmp := 0
	for tmp != slow {
		tmp = nums[tmp]
		slow = nums[slow]
	}
	return tmp
}
