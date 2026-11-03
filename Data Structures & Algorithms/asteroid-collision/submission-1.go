func abs(x int)int{
	if x < 0 {return -x}
	return x
}


func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)
	for _, asteroid := range asteroids{
		if len(stack) == 0 {
			stack = append(stack,asteroid)
		} else {
			for {
				if len(stack) == 0 {
					stack = append(stack, asteroid)
					break
				}
				top := stack[len(stack)-1]
				if (top > 0 && asteroid > 0) || (top < 0 && asteroid < 0) || (top < 0 && asteroid > 0) {
					stack = append(stack,asteroid)
					break
				} else {
					if abs(asteroid) == abs(top){
						stack = stack[:len(stack)-1]
						break
					} else if abs(asteroid) < abs(top){
						break
					} else {
						stack = stack[:len(stack)-1]
					}
				}
				
			}
		}
	}
	return stack
}
