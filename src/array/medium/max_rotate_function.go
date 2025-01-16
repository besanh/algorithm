package medium

/*
 * Input: nums = [4,3,2,6]
 * Output: 26
 * Explanation:
 * F(0) = (0 * 4) + (1 * 3) + (2 * 2) + (3 * 6) = 0 + 3 + 4 + 18 = 25
 * F(1) = (0 * 6) + (1 * 4) + (2 * 3) + (3 * 2) = 0 + 4 + 6 + 6 = 16
 * F(2) = (0 * 2) + (1 * 6) + (2 * 4) + (3 * 3) = 0 + 6 + 8 + 9 = 23
 * F(3) = (0 * 3) + (1 * 2) + (2 * 6) + (3 * 4) = 0 + 2 + 12 + 12 = 26
 * So the maximum value of F(0), F(1), F(2), F(3) is F(3) = 26.
 */
func MaxRotateFunction(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	sum := 0
	for _, item := range nums {
		sum += item
	}

	f0 := 0
	for i, item := range nums {
		f0 += i * item
	}

	maxF, currentF := f0, f0

	for i := 1; i < n; i++ {
		// Formular: F(k+1)=F(k)+sum(nums)−n×nums[n−k−1]
		currentF = currentF + sum - n*nums[n-i]
		if currentF > maxF {
			maxF = currentF
		}
	}

	return maxF
}
