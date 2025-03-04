package medium

/*
 * notme
 */
func CheckSubarraySum(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}

	if k == 0 {
		for i := 1; i < len(nums); i++ {
			if nums[i] == 0 && nums[i-1] == 0 {
				return true
			}
		}
		return false
	}

	remainderMap := map[int]int{0: -1}
	runningSum := 0

	for i, num := range nums {
		runningSum += num
		remainder := runningSum % k

		if prevIndex, exists := remainderMap[remainder]; exists {
			// Kt subarray length min la 2
			if i-prevIndex > 1 {
				return true
			}
		} else {
			remainderMap[remainder] = i
		}
	}

	return false
}
