func climbStairs(n int) int {
	mapNumOfWayClimbFrom := make(map[int]int)
	var numOfWayClimbFrom func(x int)int
	numOfWayClimbFrom = func(x int)int{
		if x < 0 {return 0}
		if x == 0 {
			return 1
		}
		if val, existed := mapNumOfWayClimbFrom[x]; existed{
			return val
		}
		result := 0
		result += numOfWayClimbFrom(x-1)
		result += numOfWayClimbFrom(x-2)
		mapNumOfWayClimbFrom[x]=result
		return result
	}
	return numOfWayClimbFrom(n)
}
