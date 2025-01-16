package easy

/*
 * Input: s = "abc", t = "ahbgdc"
 * Output: true
 * Input: s = "axc", t = "ahbgdc"
 * Output: false
 * "ace" is a subsequence of "abcde" while "aec" is not
 * Use 2 pointers i and j
 * notme
 */
func IsSubSequence(s string, t string) bool {
	m, n := len(s), len(t)
	i, j := 0, 0

	for i < m && j < n {
		if s[i] == t[j] {
			i++
		}
		// Constantly increase j
		j++
	}

	return i == m
}
