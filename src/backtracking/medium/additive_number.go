package backtracking

import "math/big"

/*
 * Input: "112358"
 * Output: true
 * Explanation:
 * The digits can form an additive sequence: 1, 1, 2, 3, 5, 8.
 * 1 + 1 = 2, 1 + 2 = 3, 2 + 3 = 5, 3 + 5 = 8
 *
 * Input: "199100199"
 * Output: true
 * Explanation:
 * The additive sequence is: 1, 99, 100, 199.
 * 1 + 99 = 100, 99 + 100 = 199
 *
 * notme
 */
func IsAdditiveNumber(num string) bool {
	n := len(num)

	var backtrack func(start int, count int, first, second *big.Int) bool
	backtrack = func(start int, count int, first, second *big.Int) bool {
		if count >= 3 && start == n {
			return true
		}

		for i := start + 1; i <= n; i++ {
			currentStr := num[start:i]

			if len(currentStr) > 1 && currentStr[0] == '0' {
				break
			}

			currentNum := new(big.Int)
			currentNum.SetString(currentStr, 10)

			if count < 2 {
				if backtrack(i, count+1, second, currentNum) {
					return true
				}
			} else {
				sum := new(big.Int).Add(first, second)
				if sum.Cmp(currentNum) != 0 {
					continue
				}

				if backtrack(i, count+1, second, currentNum) {
					return true
				}
			}
		}

		return false
	}

	return backtrack(0, 0, nil, nil)
}
