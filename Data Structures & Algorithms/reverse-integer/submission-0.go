func reverse(x int) int {
    // Check whether x is negative using its sign bit.
    negative := (uint32(x) >> 31) == 1

    // Convert to a positive number safely.
    // Use int64 here would violate the problem constraint.
    // Instead, handle the negative number directly.
    result := 0

    for x != 0 {
        digit := x % 10
        x /= 10

        // Check overflow before multiplying by 10.
        if result > 214748364 ||
            (result == 214748364 && digit > 7) {
            return 0
        }

        if result < -214748364 ||
            (result == -214748364 && digit < -8) {
            return 0
        }

        result = result*10 + digit
    }

    // The sign is already preserved by digit extraction.
    _ = negative
    return result
}