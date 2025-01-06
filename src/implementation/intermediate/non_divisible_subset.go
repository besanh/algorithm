package intermediate

/*
 * k = 3, S = [1, 7, 2, 4]
 * 1 + 7 = 8
 * 1 + 2 = 3
 * 1 + 4 = 5
 * 7 + 2 = 9
 * 7 + 4 = 11
 * 2 + 4 = 6
 * S' = {1, 7, 4}, Output = 3
 */
func NonDivisibleSubset(k int32, s []int32) int32 {
	// Write your code here
	remainderCount := make([]int32, k)
	for _, item := range s {
		// Find frequency
		remainder := item % k
		remainderCount[remainder]++
	}

	// Start with at most one number from the remainder 0 group
	result := int32(0)
	if remainderCount[0] > 0 {
		result++
	}

	// Iterate through the possible remainders up to k/2
	for i := int32(1); i <= k/2; i++ {
		if i == k-i { // Special case for remainders exactly in the middle
			if remainderCount[i] > 0 {
				result++
			}
		} else {
			// Take the maximum count of i or k-i
			result += max(remainderCount[i], remainderCount[k-i])
		}
	}

	return result
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}
