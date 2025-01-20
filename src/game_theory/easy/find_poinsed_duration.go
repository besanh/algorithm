package easy

/*
 * Formula: time = t + duration - 1
 */
func FindPoisonedDuration(timeSeries []int, duration int) int {
	if duration == 0 || len(timeSeries) == 0 {
		return 0
	}

	total := 0
	for i := 0; i < len(timeSeries)-1; i++ {
		diff := timeSeries[i+1] - timeSeries[i]
		if diff < duration {
			total += diff
		} else {
			total += duration
		}
	}

	// Add the last duration
	total += duration
	return total
}

func FindPoisonedDuration2(timeSeries []int, duration int) int {
	if duration == 0 || len(timeSeries) == 0 {
		return 0
	}

	total, end := 0, -1

	for _, t := range timeSeries {
		currentEnd := t + duration - 1
		// Khong chong lan
		if t > end {
			total += duration
		} else {
			total += (currentEnd - end)
		}

		if currentEnd > end {
			end = currentEnd
		}
	}
	return total
}
