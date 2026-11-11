type NumArray struct {
    prefixSum []int
}


func Constructor(nums []int) NumArray {
    sum := 0
	prefixSum := make([]int,0)
	for _, num := range nums{
		sum = sum + num
		prefixSum = append(prefixSum, sum)
	}
	return NumArray{
		prefixSum,
	}
}


func (this *NumArray) SumRange(left int, right int) int {
    if left == 0 {
		return this.prefixSum[right]
	}
	return this.prefixSum[right] - this.prefixSum[left-1]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */