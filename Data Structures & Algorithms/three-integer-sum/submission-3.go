
import "slices"

func threeSum(nums []int) [][]int {
    slices.Sort(nums)

    n := len(nums)
    results := make([][]int, 0)

    for i := 0; i < n-2; i++ {
        // Skip duplicate first elements.
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        // All remaining numbers are positive.
        if nums[i] > 0 {
            break
        }

        left, right := i+1, n-1

        for left < right {
            sum := nums[i] + nums[left] + nums[right]

            if sum < 0 {
                left++
            } else if sum > 0 {
                right--
            } else {
                results = append(results, []int{
                    nums[i], nums[left], nums[right],
                })

                left++
                right--

                // Skip duplicate second elements.
                for left < right && nums[left] == nums[left-1] {
                    left++
                }

                // Skip duplicate third elements.
                for left < right && nums[right] == nums[right+1] {
                    right--
                }
            }
        }
    }

    return results
}
