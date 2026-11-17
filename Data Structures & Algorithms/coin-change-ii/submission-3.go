
import "slices"

func change(amount int, coins []int) int {
    slices.Sort(coins)

    memo := make(map[[2]int]int)

    var dfs func(amount, minCoinIdx int) int
    dfs = func(amount, minCoinIdx int) int {
        if amount == 0 {
            return 1
        }
        if amount < 0 {
            return 0
        }

        key := [2]int{amount, minCoinIdx}
        if val, ok := memo[key]; ok {
            return val
        }

        ways := 0

        for i := minCoinIdx; i < len(coins); i++ {
            ways += dfs(amount-coins[i], i)
        }

        memo[key] = ways
        return ways
    }

    return dfs(amount, 0)
}
