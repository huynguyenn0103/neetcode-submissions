
func characterReplacement(s string, k int) int {
    result := 0
    n := len(s)

    for target := byte('A'); target <= 'Z'; target++ {
        left := 0
        replacements := 0

        for right := 0; right < n; right++ {
            if s[right] != target {
                replacements++
            }

            for replacements > k {
                if s[left] != target {
                    replacements--
                }
                left++
            }

            result = max(result, right-left+1)
        }
    }

    return result
}
