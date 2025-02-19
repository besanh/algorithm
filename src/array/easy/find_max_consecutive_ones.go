package easy

/*
 * Problem: https://leetcode.com/problems/max-consecutive-ones/?envType=problem-list-v2&envId=array
 * Solution: iterate each item, if item is 1, I add 1 to count, if item is 0, I reset count to 0
 * Time complexity: O(n), because I iterate only once
 * Space complexity: O(1), because I iterate each item
 * DSA: normal
 */
func FindMaxConsecutiveOnes(nums []int) int {
	rs, current := 0, 0
	for _, val := range nums {
		if val == 1 {
			current++
			if current > rs {
				rs = current
			}
		} else if val == 0 {
			current = 0
		}
	}

	return rs
}
