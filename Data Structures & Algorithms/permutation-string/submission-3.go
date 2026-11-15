
func checkInclusion(s1 string, s2 string) bool {
    m, n := len(s1), len(s2)
    if m > n {
        return false
    }

    var count [26]int

    // Count characters in s1 and subtract first window of s2.
    for i := 0; i < m; i++ {
        count[s1[i]-'a']++
        count[s2[i]-'a']--
    }

    // Check whether the first window is a permutation.
    if isZero(count) {
        return true
    }

    // Slide the window across s2.
    for i := m; i < n; i++ {
        // Remove the character leaving the window.
        count[s2[i-m]-'a']++

        // Add the character entering the window.
        count[s2[i]-'a']--

        if isZero(count) {
            return true
        }
    }

    return false
}

func isZero(count [26]int) bool {
    for _, v := range count {
        if v != 0 {
            return false
        }
    }
    return true
}
