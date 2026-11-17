import "slices"

func permute(nums []int) [][]int {
	result := make([][]int,0)
	permutation := make([]int,0)
	visited := make(map[int]bool)
	n := len(nums)
	var buildPermutation func()
	buildPermutation = func(){
		if len(permutation) == n{
			result = append(result, slices.Clone(permutation))
			return
		} 
		for i := 0; i < n; i++ {
			if visited[nums[i]]{continue}
			visited[nums[i]]=true
			permutation = append(permutation, nums[i])
			buildPermutation()
			visited[nums[i]]=false
			permutation = permutation[:len(permutation)-1]
		}
	}
	buildPermutation()
	return result
}
