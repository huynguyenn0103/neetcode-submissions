//2 1

// 10 4.5 3 3 4 4

import "slices"

type Car struct{
	position int
	speed int
}

func timeSpend(car Car, target int) float64{
	return float64(target-car.position)/float64(car.speed)
}

func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	cars := make([]Car, 0)
	for i := 0; i < n; i++ {
		cars = append(cars, Car{position: position[i], speed: speed[i]})
	}
	slices.SortFunc(cars, func(c1,c2 Car)int{return c1.position-c2.position})
	stack := make([]float64,0)
	for _, car := range cars{
		timeNeed := timeSpend(car, target)
		for len(stack) > 0 && timeNeed >= stack[len(stack)-1]{
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, timeNeed)
	}
	return len(stack)
}
