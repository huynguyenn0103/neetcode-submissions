//1 : 3 2 1
//2 : 2 1 2
//3 : 1 2 1
//4 : 2 1
//5 : 1 1
//6 : 0

func timeRequiredToBuy(tickets []int, k int) int {
    time := 0
	target := k
	for{
		time++
		first := tickets[0]
		tickets = tickets[1:]
		first--
		if first == 0{
			if target == 0{
				return time
			}
		} else {
			tickets = append(tickets, first)
		}
		target--
		if target < 0 {
			target = len(tickets)-1
		}
	}
	return 0
}