func maxProduct(nums []int) int {
    n := len(nums)
    cache := make(map[int][2]int)

    var dp func(i int) (int, int)

    dp = func(i int) (int, int) {
        if i == 0 {
            return nums[0], nums[0]
        }

        if cached, ok := cache[i]; ok {
            return cached[0], cached[1]
        }

        prevMax, prevMin := dp(i - 1)

        num := nums[i]

        maxProduct := max(
            num,
            max(num*prevMax, num*prevMin),
        )

        minProduct := min(
            num,
            min(num*prevMax, num*prevMin),
        )

        cache[i] = [2]int{maxProduct, minProduct}

        return maxProduct, minProduct
    }

    result := nums[0]

    for i := 0; i < n; i++ {
        maxProduct, _ := dp(i)
        result = max(result, maxProduct)
    }

    return result
}