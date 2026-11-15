
func canCompleteCircuit(gas []int, cost []int) int {
    totalGas := 0
    totalCost := 0

    start := 0
    tank := 0

    for i := 0; i < len(gas); i++ {
        diff := gas[i] - cost[i]

        totalGas += gas[i]
        totalCost += cost[i]

        tank += diff

        if tank < 0 {
            // Current starting point cannot work.
            start = i + 1
            tank = 0
        }
    }

    if totalGas < totalCost {
        return -1
    }

    return start
}
