type TrieNode struct {
    children [26]*TrieNode
    word     string
}

func findWords(board [][]byte, words []string) []string {
    if len(board) == 0 || len(board[0]) == 0 {
        return nil
    }

    m, n := len(board), len(board[0])

    // Step 1: Build Trie
    root := &TrieNode{}

    for _, word := range words {
        node := root

        for i := 0; i < len(word); i++ {
            idx := word[i] - 'a'

            if node.children[idx] == nil {
                node.children[idx] = &TrieNode{}
            }

            node = node.children[idx]
        }

        node.word = word
    }

    result := []string{}

    directions := [][2]int{
        {-1, 0},
        {0, 1},
        {1, 0},
        {0, -1},
    }

    var dfs func(row, col int, node *TrieNode)

    dfs = func(row, col int, node *TrieNode) {
        if row < 0 || row >= m || col < 0 || col >= n {
            return
        }

        char := board[row][col]

        // Already visited
        if char == '#' {
            return
        }

        // Current prefix doesn't exist in Trie
        idx := char - 'a'
        next := node.children[idx]

        if next == nil {
            return
        }

        // Found a complete word
        if next.word != "" {
            result = append(result, next.word)

            // Avoid returning the same word multiple times
            next.word = ""
        }

        // Mark cell as visited
        board[row][col] = '#'

        for _, dir := range directions {
            dfs(row+dir[0], col+dir[1], next)
        }

        // Restore cell after backtracking
        board[row][col] = char
    }

    // Step 2: Search from every cell
    for row := 0; row < m; row++ {
        for col := 0; col < n; col++ {
            dfs(row, col, root)
        }
    }

    return result
}