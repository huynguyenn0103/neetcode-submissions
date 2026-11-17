
import "slices"

func subsetsWithDup(nums []int) [][]int {
    slices.Sort(nums)

    result := make([][]int, 0)
    subset := make([]int, 0)

    var backtrack func(start int)
    backtrack = func(start int) {
        result = append(result, slices.Clone(subset))

        for i := start; i < len(nums); i++ {
            // Skip duplicates at the same recursion level.
            if i > start && nums[i] == nums[i-1] {
                continue
            }

            subset = append(subset, nums[i])
            backtrack(i + 1)
            subset = subset[:len(subset)-1]
        }
    }

    backtrack(0)
    return result
}
