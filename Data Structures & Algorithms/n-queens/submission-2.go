
func solveNQueens(n int) [][]string {
    results := make([][]string, 0)

    // Store the column of the queen in each row.
    queens := make([]int, n)

    var backtrack func(row int)
    backtrack = func(row int) {
        // All rows have a queen.
        if row == n {
            board := make([]string, n)

            for i := 0; i < n; i++ {
                line := make([]byte, n)

                for j := 0; j < n; j++ {
                    line[j] = '.'
                }

                line[queens[i]] = 'Q'
                board[i] = string(line)
            }

            results = append(results, board)
            return
        }

        // Try every column in the current row.
        for col := 0; col < n; col++ {
            valid := true

            // Check previously placed queens.
            for prevRow := 0; prevRow < row; prevRow++ {
                prevCol := queens[prevRow]

                // Same column or same diagonal?
                if prevCol == col ||
                    abs(prevCol-col) == abs(prevRow-row) {
                    valid = false
                    break
                }
            }

            if !valid {
                continue
            }

            // Choose
            queens[row] = col

            // Explore
            backtrack(row + 1)

            // Undo
            queens[row] = -1
        }
    }

    backtrack(0)
    return results
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
