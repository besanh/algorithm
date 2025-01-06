package week1

/*
 * Using slide window
 * @param d int32: number of days
 * @param m int32: number of months => length of subarray
 * s = [2, 2, 1, 3, 2], d = 4, m = 2
 * Calculate: 2+2 = 4, 1+3 = 4
 * Output: 2
 */
func Birthday(s []int32, d int32, m int32) int32 {
	// Write your code here
	var window, sum int32

	if m > int32(len(s)) {
		return 0
	}

	// Calculate initial window
	for i := 0; i < int(m); i++ {
		window += s[i]
	}

	// Check if window matches d
	if window == d {
		sum++
	}

	// Move window from m to len(s)
	for i := m; i < int32(len(s)-1); i++ {
		// Remove first element of window and add new element
		window += s[i] - s[i-m]
		if window == d {
			sum++
		}
	}

	return sum
}
