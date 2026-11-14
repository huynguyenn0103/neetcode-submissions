func minSwaps(s string) int {
    balance := 0
    maxNegative := 0

    for i := 0; i < len(s); i++ {
        if s[i] == '[' {
            balance++
        } else {
            balance--
        }

        if balance < 0 {
            maxNegative = max(maxNegative, -balance)
        }
    }

    return (maxNegative + 1) / 2
}