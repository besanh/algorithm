package medium

/*
 * Input: s = "abpcplea", dictionary = ["ale","apple","monkey","plea"]
 * Output: "apple"
 * Input: s = "abpcplea", dictionary = ["a","b","c"]
 * Output: "a"
 */
func FindLongestWord(s string, dictionary []string) string {
	best := ""

	for _, word := range dictionary {
		if isSubSequence(s, word) {
			// If word is longer than best or word is equal to best but word is lexicographically smaller than best
			if len(word) > len(best) || (len(word) == len(best) && word < best) {
				best = word
			}
		}
	}

	return best
}

// Use 2 pointer to check sub sequence exist in s
func isSubSequence(s, sub string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(sub) {
		if s[i] == sub[j] {
			j++
		}
		i++
	}
	return j == len(sub)
}
