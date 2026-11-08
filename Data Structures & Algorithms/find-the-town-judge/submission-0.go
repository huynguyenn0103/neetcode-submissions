func findJudge(n int, trust [][]int) int {
	indegrees := make([]int,n)
	outdegrees := make([]int,n)
	for _, t := range trust {
		from, to := t[0]-1,t[1]-1
		indegrees[to]++
		outdegrees[from]++
	}
	for i := 0; i < n; i++{
		if indegrees[i] == n-1 && outdegrees[i] == 0 {
			return i+1
		}
	}
	return -1
}
