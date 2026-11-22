
func findWords(board [][]byte, words []string) []string {
    m, n := len(board), len(board[0])

    // 1. Build Trie
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

    // 2. Backtracking
    result := make([]string, 0)

    var dfs func(row, col int, node *TrieNode)

    dfs = func(row, col int, node *TrieNode) {
        if row < 0 || row >= m || col < 0 || col >= n {
            return
        }

        c := board[row][col]

        // Already visited
        if c == '#' {
            return
        }

        next := node.children[c-'a']

        // No word has this prefix
        if next == nil {
            return
        }

        // Found a complete word
        if next.word != "" {
            result = append(result, next.word)
            next.word = "" // Avoid duplicate results
        }

        // Mark current cell as visited
        board[row][col] = '#'

        dfs(row+1, col, next)
        dfs(row-1, col, next)
        dfs(row, col+1, next)
        dfs(row, col-1, next)

        // Restore board
        board[row][col] = c
    }

    // 3. Start DFS from every cell
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            dfs(i, j, root)
        }
    }

    return result
}

type TrieNode struct {
    children [26]*TrieNode
    word     string
}
