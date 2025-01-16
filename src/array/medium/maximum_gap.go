package medium

import (
	"sort"
)

/*
 * Input: nums = [3,6,9,1]
 * Output: 3
 * Explanation: The sorted form of the array is [1,3,6,9], either (3,6) or (6,9) has the maximum difference 3.
 * Methodology: buble sort, 2 pointers, greedy
 */
func MaximumGap(nums []int) int {
	sort.Ints(nums)
	slow, fast := 0, 1
	rs := 0
	for fast < len(nums) {
		rs = max(nums[fast]-nums[slow], rs)
		slow++
		fast++
	}

	return rs
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
