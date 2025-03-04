package medium

/*
 * Input: nums = [1,1,1,1,1], target = 3
 * Output: 5
 * Explanation: There are 5 ways to assign symbols to make the sum of nums be target 3.
 * -1 + 1 + 1 + 1 + 1 = 3
 * +1 - 1 + 1 + 1 + 1 = 3
 * +1 + 1 - 1 + 1 + 1 = 3
 * +1 + 1 + 1 - 1 + 1 = 3
 * +1 + 1 + 1 + 1 - 1 = 3
 *
 * Input: nums = [1], target = 1
 * Output: 1
 */
func FindTargetSumWays(nums []int, target int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	if total < target || (total+target)%2 != 0 {
		return 0
	}

	sum := (total + target) / 2
	if sum < 0 {
		return 0
	}

	dp := make([]int, sum+1)
	// Ensure at least 1 way to reach sum
	dp[0] = 1
	for _, val := range nums {
		for j := sum; j >= val; j-- {
			dp[j] += dp[j-val]
		}
	}

	return dp[sum]
}
