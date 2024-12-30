package week2

/*
 * Game theory
 *
 *
 */
func TowerBreakers(n int32, m int32) int32 {
	// Write your code here
	// If m == 1, player 2 wins, because the initial tall is both 1
	if m == 1 {
		return 2
	}

	// If n is even, player 2 wins
	if n%2 == 0 {
		return 2
	}
	return 1
}
