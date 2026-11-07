import "slices"

func makesquare(matchsticks []int) bool {
    n := len(matchsticks)
    if n < 4 {
        return false
    }

    sum := 0
    for _, v := range matchsticks {
        sum += v
    }

    if sum%4 != 0 {
        return false
    }

    target := sum / 4

    // Largest matchsticks first to improve pruning.
    slices.Sort(matchsticks)
    slices.Reverse(matchsticks)

    // If the largest stick exceeds target, impossible.
    if matchsticks[0] > target {
        return false
    }

    sides := [4]int{}

    var backtrack func(idx int) bool
    backtrack = func(idx int) bool {
        // All matchsticks have been assigned.
        if idx == n {
            return true
        }

        for i := 0; i < 4; i++ {
            // Cannot exceed the target.
            if sides[i]+matchsticks[idx] > target {
                continue
            }

            sides[i] += matchsticks[idx]

            if backtrack(idx + 1) {
                return true
            }

            sides[i] -= matchsticks[idx]

            // If this side was empty, other empty sides
            // are symmetric, so don't try them.
            if sides[i] == 0 {
                break
            }
        }

        return false
    }

    return backtrack(0)
}