package medium

import "sort"

/*
 * Using backtracking
 * Input: nums = [1,1,2]
 * Output:
 * [[1,1,2],
 * [1,2,1],
 * [2,1,1]]
 */
func PermuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	rs := make([][]int, 0)
	current := make([]int, 0)
	used := make([]bool, len(nums))

	var backtrack func()
	backtrack = func() {
		if len(current) == len(nums) {
			perm := make([]int, len(current))
			copy(perm, current)
			rs = append(rs, perm)

			return
		}

		for i := 0; i < len(nums); i++ {
			// Bỏ qua phần tử đã được sử dụng
			if used[i] {
				continue
			}
			// Bỏ qua phần tử trùng lặp nếu phần tử trước chưa được sử dụng
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}

			// Chọn phần tử nums[i]
			used[i] = true
			current = append(current, nums[i])

			// Đệ quy để xây dựng hoán vị tiếp theo
			backtrack()

			// Bỏ chọn phần tử nums[i]
			used[i] = false
			current = current[:len(current)-1]
		}
	}

	backtrack()

	return rs
}
