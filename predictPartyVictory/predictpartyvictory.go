package main

func predictPartyVictory(senate string) string {
	radiantQueue := make([]int, 0)
	direQueue := make([]int, 0)

	// Initialize the queues with indices of senators
	for i, s := range senate {
		if s == 'R' {
			radiantQueue = append(radiantQueue, i)
		} else {
			direQueue = append(direQueue, i)
		}
	}

	for len(radiantQueue) > 0 && len(direQueue) > 0 {
		radiantIdx := radiantQueue[0]
		direIdx := direQueue[0]

		// The senator with the smallest index exercises their right first
		if radiantIdx < direIdx {
			radiantQueue = append(radiantQueue, radiantIdx+len(senate))
		} else {
			direQueue = append(direQueue, direIdx+len(senate))
		}

		radiantQueue = radiantQueue[1:]
		direQueue = direQueue[1:]
	}

	if len(radiantQueue) > 0 {
		return "Radiant"
	}

	return "Dire"
}
