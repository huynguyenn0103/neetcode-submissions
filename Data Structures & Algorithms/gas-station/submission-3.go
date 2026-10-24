func canCompleteCircuit(gas []int, cost []int) int {
	candidateIdx := 0
	n := len(gas)
	for {
		if candidateIdx == n {return -1}
		if gas[candidateIdx] >= cost[candidateIdx]{
			curIdx := candidateIdx
			curGas := gas[candidateIdx]
			for {
				if curGas < cost[curIdx]{
					if curIdx <= candidateIdx{return -1}
					candidateIdx = curIdx
					break
				}
				curGas = curGas-cost[curIdx]
				curIdx++
				if curIdx == n {
					curIdx = 0
				}
				curGas = curGas + gas[curIdx]
				if curIdx == candidateIdx{
					return candidateIdx
				}
			}
		} else {
			candidateIdx++
		}
	}
	return -1
}
