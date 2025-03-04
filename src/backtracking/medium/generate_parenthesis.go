package backtracking

/*
 * Input: n = 3
 * Output: ["((()))","(()())","(())()","()(())","()()()"]
 */
func GenerateParenthesis(n int) []string {
	rs := make([]string, 0)

	var backtrack func(current string, open, close int)
	backtrack = func(current string, open, close int) {
		// length of item = 2*n
		if len(current) == 2*n {
			rs = append(rs, current)
			return
		}

		if open < n {
			backtrack(current+"(", open+1, close)
		}

		if close < open {
			backtrack(current+")", open, close+1)
		}
	}

	backtrack("", 0, 0)

	return rs
}
