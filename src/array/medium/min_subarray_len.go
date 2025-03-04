package medium

/*
 * Problem: https://leetcode.com/problems/minimum-size-subarray-sum/?envType=problem-list-v2&envId=array
 * Reference: https://app.diagrams.net/#G1eKRONRBBid0_9h5cvYrpOvIJu5kBQ9zJ#%7B%22pageId%22%3A%22LCAZqEOmr3nNQPfN1sSf%22%7D
 * Approach: use sliding window
 * Time complexity: O(n)
 * Space complexity: O(1)
 * DSA: Sliding Window
 */
func MinSubArrayLen(target int, nums []int) int {
	sum, left := 0, 0
	n := len(nums)

	// Initialize minLen to n+1 to make sure that minLen will be updated
	// if there is no valid subarray
	minLen := n + 1

	for right := 0; right < n; right++ {
		sum += nums[right]

		for sum >= target {
			currentLen := right - left + 1
			if currentLen < minLen {
				minLen = currentLen
			}
			sum -= nums[left]
			left++
		}
	}

	if minLen == n+1 {
		return 0
	}

	return minLen
}
