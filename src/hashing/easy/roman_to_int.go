package easy

/*
 * Problem: https://leetcode.com/problems/roman-to-integer/description/?envType=problem-list-v2&envId=hash-table
 * DSA: Hash Table
 * Time complexity: O(1)
 * Space complexity: O(1)
 */
func RomanToInt(s string) int {
	hash := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result, prev := 0, 0

	for i := len(s) - 1; i >= 0; i-- {
		current := hash[s[i]]
		if current < prev {
			result -= current
		} else {
			result += current
		}
		prev = current
	}

	return result
}
