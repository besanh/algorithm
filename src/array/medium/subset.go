package medium

/*
 * Input: [1, 2, 3]
 * Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
 * Using backtracking or recursive
 */
func Subset(nums []int) [][]int {
	rs := make([][]int, 0)
	current := []int{}

	var backtracking func(start int)
	backtracking = func(start int) {
		tmp := make([]int, len(current))
		copy(tmp, current)
		rs = append(rs, tmp)

		for i := start; i < len(nums); i++ {
			current = append(current, nums[i])
			backtracking(i + 1)
			current = current[:len(current)-1]
		}
	}

	backtracking(0)

	return rs
}
