func buildSumSquare(n int)int{
	sum := 0
	for n != 0 {
		remainder := n % 10
		sum += remainder * remainder
		n=n/10
	}
	return sum
}

func isHappy(n int) bool {
    visited := make(map[int]bool)
	visited[n] = true
	val := n
	for val != 1 {
		val = buildSumSquare(val)
		if visited[val]{return false}
		visited[val] = true
	}
	return true
}
