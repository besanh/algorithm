package backtracking

/*
 * n = 4, k = 2
 * Output: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
 */
func Combine(n int, k int) [][]int {
	rs := make([][]int, 0)
	combination := make([]int, 0)

	var backtracking func(start int)
	backtracking = func(start int) {
		if len(combination) == k {
			tmp := make([]int, k)
			copy(tmp, combination)
			rs = append(rs, tmp)
			return
		}

		for i := start; i <= n; i++ {
			combination = append(combination, i)
			backtracking(i + 1)
			combination = combination[:len(combination)-1]
		}
	}

	backtracking(1)
	return rs
}
