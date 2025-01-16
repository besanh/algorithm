package backtracking

import "sort"

/*
 * Input: nums = [1,2,2]
 * Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]
 */
func SubsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	rs := make([][]int, 0)
	combination := make([]int, 0)

	var backtracking func(start int)
	backtracking = func(start int) {
		tmp := make([]int, len(combination))
		copy(tmp, combination)
		rs = append(rs, tmp)

		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			combination = append(combination, nums[i])
			backtracking(i + 1)
			combination = combination[:len(combination)-1]
		}
	}

	backtracking(0)
	return rs
}
