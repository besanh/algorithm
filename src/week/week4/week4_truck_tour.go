package week4

/*
 * [][]int32{{1, 5}, {10, 3}, {3, 4}}
 * Output: 1
 */
func TruckTour(petrolpumps [][]int32) int32 {
	// Write your code here
	var amountPetrol, totalDistance, tank, start int32
	for _, val := range petrolpumps {
		amountPetrol += val[0]
		totalDistance += val[1]
	}

	for i, val := range petrolpumps {
		tank += val[0] - val[1]
		if tank < 0 {
			tank = 0
			start = int32(i + 1)
		}
	}

	if amountPetrol >= totalDistance {
		return start
	}

	return -1
}

// leetcode
func CanCompleteCircuit(gas []int, cost []int) int {
	start, totalTank, currentTank := 0, 0, 0

	for i, item := range gas {
		totalTank += item - cost[i]
		currentTank += item - cost[i]

		if currentTank < 0 {
			start = i + 1
			currentTank = 0
		}
	}

	if totalTank < 0 {
		return -1
	}

	return start
}
