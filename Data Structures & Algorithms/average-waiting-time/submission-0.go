func averageWaitingTime(customers [][]int) float64 {
    curTime := customers[0][0]
	waitingTime := 0
	for _, customer := range customers{
		arrivalTime, processTime := customer[0], customer[1]
		if arrivalTime > curTime {
			curTime = arrivalTime
		} else {
			waitingTime += curTime - arrivalTime
		}

		curTime += processTime
		waitingTime += processTime
	}
	return float64(waitingTime) / float64(len(customers))
}