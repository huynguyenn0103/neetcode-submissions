func sumOfSquare(n int) int {
	total := 0
	for n != 0 {
		x := n % 10
		total += x*x
		n = n/10
	}
	return total
}
func isHappy(n int) bool {
	total := n
	mapVisited := make(map[int]bool)
	mapVisited[n]=true
    for {
		total = sumOfSquare(total)
		if total == 1{
			return true
		}
		if mapVisited[total] {
			return false
		}
		mapVisited[total]=true
	}
	return false
}
