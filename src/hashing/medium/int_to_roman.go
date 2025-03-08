package medium

/*
 * Problem: https://leetcode.com/problems/integer-to-roman/submissions/1566655707/?envType=problem-list-v2&envId=hash-table
 * Solution: delcare needed hash table, convert number to roman
 * Time complexity: O(1)
 * Space complexity: O(1)
 * DSA: Hash Table
 */
func IntToRoman(num int) string {
	hash := []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var result string

	for _, h := range hash {
		for num >= h.value {
			result += h.symbol
			num -= h.value
		}
	}

	return result
}
