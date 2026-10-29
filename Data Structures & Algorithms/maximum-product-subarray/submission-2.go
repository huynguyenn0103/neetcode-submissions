func maxProduct(nums []int) int {
    maxProd := nums[0]
    minProd := nums[0]
    result := nums[0]

    for i := 1; i < len(nums); i++ {
        num := nums[i]

        a := num
        b := num * maxProd
        c := num * minProd

        maxProd = max(a, b, c)
        minProd = min(a, b, c)

        result = max(result, maxProd)
    }

    return result
}